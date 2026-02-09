package main

import (
	"admin/internal/config"
	"admin/internal/handler"
	"admin/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 添加静态文件服务
	if c.Static.Dir != "" {
		// 获取绝对路径
		absDir, err := filepath.Abs(c.Static.Dir)
		if err != nil {
			fmt.Printf("Error getting absolute path: %v\n", err)
			return
		}
		// 检查目录是否存在
		if _, err := os.Stat(absDir); os.IsNotExist(err) {
			fmt.Printf("Static directory does not exist: %s\n", absDir)
			return
		}
		// 使用标准库的http.FileServer
		fileServer := http.FileServer(http.Dir(absDir))
		// 创建包装函数来处理路径
		staticHandler := func(w http.ResponseWriter, r *http.Request) {
			// 调整请求路径，去掉前缀
			r.URL.Path = strings.TrimPrefix(r.URL.Path, c.Static.Prefix)
			if r.URL.Path == "" {
				r.URL.Path = "/"
			}
			fileServer.ServeHTTP(w, r)
		}

		// 注册静态文件路由
		server.AddRoute(rest.Route{
			Method:  http.MethodGet,
			Path:    "/common/:any",
			Handler: http.HandlerFunc(staticHandler),
		})
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
