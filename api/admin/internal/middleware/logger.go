package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	logclient "admin/grpc-client/apilog"
	"admin/internal/config"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/jwt"
	"github.com/dwrui/go-zero-admin/pkg/utils/plugin"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"
	"github.com/zeromicro/go-zero/core/logx"
)

// CustomLogger 自定义日志中间件
type CustomLogger struct {
	conf         config.Config
	logClient    logclient.LogServiceClient
	logChan      chan *LogData
	excludePaths map[string]bool
	workerWg     sync.WaitGroup
	ctx          context.Context
	cancel       context.CancelFunc
}

// LogData 日志数据结构
type LogData struct {
	UserID      int64     `json:"user_id"`      //用户ID
	AccountID   int64     `json:"account_id"`   //账号ID
	BusinessID  int64     `json:"business_id"`  //业务ID
	Type        string    `json:"type"`         //日志类型 admin后台日志 adminpro总后台日志
	Method      string    `json:"method"`       //请求方法
	Path        string    `json:"path"`         //请求路径
	IP          string    `json:"ip"`           //请求IP
	Address     string    `json:"address"`      //根据ip获取的地址
	ReqHeaders  string    `json:"req_headers"`  //请求头
	ReqBody     string    `json:"req_body"`     //请求体
	RespHeaders string    `json:"resp_headers"` //响应头
	RespBody    string    `json:"resp_body"`    //响应体
	Status      int       `json:"status"`       //1成功0失败
	Duration    int64     `json:"duration"`     // 耗时
	CreatedTime time.Time `json:"created_time"` // 创建时间
}

// NewCustomLogger 创建自定义日志中间件
func NewCustomLogger(conf config.Config, logClient logclient.LogServiceClient) *CustomLogger {
	ctx, cancel := context.WithCancel(context.Background())

	// 从配置中读取需要排除的接口路径
	excludePaths := make(map[string]bool)
	for _, path := range conf.LogExcludePaths {
		excludePaths[path] = true
	}

	logger := &CustomLogger{
		conf:         conf,
		logClient:    logClient,
		logChan:      make(chan *LogData, 1000), // 缓冲通道
		excludePaths: excludePaths,
		ctx:          ctx,
		cancel:       cancel,
	}

	// 启动异步日志处理工作池
	logger.startLogWorkers(3) // 启动3个工作协程

	return logger
}

// Middleware 日志中间件处理函数
func (l *CustomLogger) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// 检查是否需要排除记录
		if l.shouldExclude(r.URL.Path) {
			next(w, r)
			return
		}

		// 包装响应写入器以捕获响应数据
		wrappedWriter := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           make([]byte, 0),
		}

		// 读取请求体
		var requestBody string
		if r.Body != nil {
			bodyBytes := make([]byte, 0, 1024)
			buf := make([]byte, 1024)
			for {
				n, err := r.Body.Read(buf)
				if n > 0 {
					bodyBytes = append(bodyBytes, buf[:n]...)
				}
				if err != nil {
					break
				}
			}
			if len(bodyBytes) > 0 && len(bodyBytes) < 10240 { // 限制请求体大小为10KB
				requestBody = string(bodyBytes)
			}
			// 重新设置请求体，以便后续处理
			r.Body = http.NoBody
			if len(bodyBytes) > 0 {
				r.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))
			}
		}

		// 执行下一个处理函数
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Sprintf("panic: %v", err)
					wrappedWriter.statusCode = http.StatusInternalServerError
				}
			}()
			next(wrappedWriter, r)
		}()

		// 计算处理时间
		duration := time.Since(startTime).Milliseconds()

		// 获取用户信息（如果有token）
		var userID, accountID, businessID int64
		if token := r.Header.Get("Authorization"); token != "" {
			// 这里可以根据token解析用户信息
			// 简化处理，实际需要调用认证服务
			jwtSecret := jwt.JwtConfig{
				AccessSecret: l.conf.Auth.AccessSecret,
				AccessExpire: l.conf.Auth.AccessExpire,
			}
			userInfo, err := jwt.ParseToken(jwtSecret, token)
			if err != nil {
				logx.Errorf("ParseToken error: %v", err)
			}
			if userInfo != nil {
				userID = ga.Int64(userInfo.UserId)
				accountID = ga.Int64(userInfo.UserId)
				businessID = ga.Int64(userInfo.BusinessId)
			}
		}
		fmt.Println(ga.GetIp(r))
		address, err := plugin.NewIpRegion(ga.GetIp(r))
		if err != nil {
			address = ""
		}
		req_str, _ := json.Marshal(r.Header)
		rep_str, _ := json.Marshal(wrappedWriter.Header())
		// 构建日志数据
		logData := &LogData{
			UserID:      userID,
			AccountID:   accountID,
			BusinessID:  businessID,
			Type:        "admin",
			Method:      r.Method,
			Path:        r.URL.Path,
			IP:          ga.GetIp(r),
			Address:     address,
			ReqHeaders:  ga.String(req_str),
			ReqBody:     requestBody,
			RespHeaders: ga.String(rep_str),
			RespBody:    string(wrappedWriter.body),
			Status:      wrappedWriter.statusCode,
			Duration:    duration,
			CreatedTime: time.Now(),
		}
		// 异步发送到日志通道
		select {
		case l.logChan <- logData:
		default:
			// 通道满了，记录警告
			logx.Errorf("Log channel is full, dropping log: %s %s", r.Method, r.URL.Path)
		}

		// 将响应写回原始ResponseWriter
		for key, values := range wrappedWriter.Header() {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(wrappedWriter.statusCode)
		if len(wrappedWriter.body) > 0 {
			w.Write(wrappedWriter.body)
		}
	}
}

// shouldExclude 检查是否需要排除记录
func (l *CustomLogger) shouldExclude(path string) bool {
	// 精确匹配检查
	if l.excludePaths[path] {
		return true
	}

	// 前缀匹配检查
	excludePrefixes := []string{
		"/static/",
		"/assets/",
		"/favicon.ico",
		"/health",
		"/metrics",
	}

	for _, prefix := range excludePrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}

	return false
}

// responseWriter 包装ResponseWriter以捕获响应数据
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.body = append(w.body, data...)
	return len(data), nil
}

// UserInfo 用户信息
type UserInfo struct {
	UserID     int64
	AccountID  int64
	BusinessID int64
}

// startLogWorkers 启动日志处理工作池
func (l *CustomLogger) startLogWorkers(workerCount int) {
	for i := 0; i < workerCount; i++ {
		l.workerWg.Add(1)
		go l.logWorker(i)
	}
}

// logWorker 日志处理工作协程
func (l *CustomLogger) logWorker(id int) {
	defer l.workerWg.Done()

	ticker := time.NewTicker(5 * time.Second) // 每5秒批量处理一次
	defer ticker.Stop()

	var batch []*LogData

	for {
		select {
		case logData, ok := <-l.logChan:
			if !ok {
				// 通道关闭，处理剩余日志
				if len(batch) > 0 {
					l.processLogBatch(batch)
				}
				return
			}
			batch = append(batch, logData)

			// 批量达到100条立即处理
			if len(batch) >= 100 {
				l.processLogBatch(batch)
				batch = batch[:0] // 清空batch
			}

		case <-ticker.C:
			// 定时处理
			if len(batch) > 0 {
				l.processLogBatch(batch)
				batch = batch[:0] // 清空batch
			}

		case <-l.ctx.Done():
			// 服务关闭
			if len(batch) > 0 {
				l.processLogBatch(batch)
			}
			return
		}
	}
}

// processLogBatch 批量处理日志
func (l *CustomLogger) processLogBatch(batch []*LogData) {
	if len(batch) == 0 {
		return
	}
	// 批量调用Log RPC服务
	for _, logData := range batch {

		// 转换为Log RPC请求
		req := &logclient.OperationLogRequest{
			UserId:      logData.UserID,
			AccountId:   logData.AccountID,
			BusinessId:  logData.BusinessID,
			Type:        logData.Type,
			Method:      logData.Method,
			Path:        logData.Path,
			Ip:          logData.IP,
			Address:     logData.Address,
			ReqHeaders:  logData.ReqHeaders,
			ReqBody:     logData.ReqBody,
			RespHeaders: logData.RespHeaders,
			RespBody:    logData.RespBody,
			Status:      int32(logData.Status),
			Duration:    logData.Duration,
		}

		// 调用Log RPC服务
		_, err := l.logClient.AddOperationLog(l.ctx, req)
		fmt.Println(logData.Path)
		fmt.Println(err)
		if err != nil {
			logx.Errorf("Failed to add operation log: %v", err)
		}
	}

	logx.Infof("Successfully processed %d logs", len(batch))
}

// Stop 停止日志中间件
func (l *CustomLogger) Stop() {
	close(l.logChan)
	l.cancel()
	l.workerWg.Wait()
}
