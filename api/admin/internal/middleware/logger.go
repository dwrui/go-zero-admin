package middleware

import (
	"admin/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// CustomLogger 自定义日志中间件
type CustomLogger struct {
	conf         config.Config
	redisClient  *redis.Redis
	logChan      chan *LogData
	excludePaths map[string]bool
	workerWg     sync.WaitGroup
	ctx          context.Context
	cancel       context.CancelFunc
}

// LogData 日志数据结构
type LogData struct {
	UserID      int64     `json:"user_id"`
	AccountID   int64     `json:"account_id"`
	BusinessID  int64     `json:"business_id"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	IP          string    `json:"ip"`
	UserAgent   string    `json:"user_agent"`
	Status      int       `json:"status"`
	Duration    int64     `json:"duration"` // 毫秒
	RequestBody string    `json:"request_body"`
	Response    string    `json:"response"`
	Error       string    `json:"error"`
	CreatedTime time.Time `json:"created_time"`
}

// NewCustomLogger 创建自定义日志中间件
func NewCustomLogger(conf config.Config, redisClient *redis.Redis) *CustomLogger {
	ctx, cancel := context.WithCancel(context.Background())

	// 需要排除的接口路径
	excludePaths := map[string]bool{
		"/v1/common/getCaptcha": true, // 验证码接口不记录
		"/v1/common/getMenu":    true, // 菜单接口
		"/v1/common/getQuick":   true, // 快捷方式接口
	}

	logger := &CustomLogger{
		conf:         conf,
		redisClient:  redisClient,
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
		var errMsg string
		func() {
			defer func() {
				if err := recover(); err != nil {
					errMsg = fmt.Sprintf("panic: %v", err)
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
			userInfo := l.parseToken(token)
			if userInfo != nil {
				userID = userInfo.UserID
				accountID = userInfo.AccountID
				businessID = userInfo.BusinessID
			}
		}

		// 构建日志数据
		logData := &LogData{
			UserID:      userID,
			AccountID:   accountID,
			BusinessID:  businessID,
			Method:      r.Method,
			Path:        r.URL.Path,
			IP:          l.getClientIP(r),
			UserAgent:   r.UserAgent(),
			Status:      wrappedWriter.statusCode,
			Duration:    duration,
			RequestBody: requestBody,
			Response:    string(wrappedWriter.body),
			Error:       errMsg,
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

// getClientIP 获取客户端IP
func (l *CustomLogger) getClientIP(r *http.Request) string {
	// 优先检查X-Forwarded-For
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// 检查X-Real-IP
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return xri
	}

	// 最后使用RemoteAddr
	ip := r.RemoteAddr
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	return ip
}

// UserInfo 用户信息
type UserInfo struct {
	UserID     int64
	AccountID  int64
	BusinessID int64
}

// parseToken 解析token获取用户信息（简化版）
func (l *CustomLogger) parseToken(token string) *UserInfo {
	// 这里应该调用认证服务解析token
	// 简化处理，返回模拟数据
	if strings.HasPrefix(token, "Bearer ") {
		// 实际应该调用认证服务
		return &UserInfo{
			UserID:     1,
			AccountID:  1,
			BusinessID: 1,
		}
	}
	return nil
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

	// 将日志数据转换为JSON
	logJSON, err := json.Marshal(batch)
	if err != nil {
		logx.Errorf("Failed to marshal log batch: %v", err)
		return
	}

	// 使用Redis队列异步存储（或调用日志服务）
	key := fmt.Sprintf("api:logs:%s", time.Now().Format("20060102"))

	// 将日志推送到Redis列表
	_, err = l.redisClient.Lpush(key, string(logJSON))
	if err != nil {
		logx.Errorf("Failed to push logs to Redis: %v", err)
		return
	}

	// 设置过期时间（7天）
	_ = l.redisClient.Expire(key, 7*24*3600)

	logx.Infof("Successfully processed %d logs", len(batch))
}

// Stop 停止日志中间件
func (l *CustomLogger) Stop() {
	close(l.logChan)
	l.cancel()
	l.workerWg.Wait()
}
