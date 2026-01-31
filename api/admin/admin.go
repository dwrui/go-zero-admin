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
	// 添加静态文件服务 - 正确的方法
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

		fmt.Printf("Static file server enabled: /common/static/ -> %s\n", absDir)

		// 创建静态文件处理函数
		staticHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 从URL中提取文件路径
			path := strings.TrimPrefix(r.URL.Path, "/common/static/")

			fmt.Printf("Static file request: %s -> %s\n", r.URL.Path, path)

			// 如果路径为空，返回404
			if path == "" || path == "/" {
				http.Error(w, "File not specified", http.StatusNotFound)
				return
			}

			// 构建完整文件路径
			fullPath := filepath.Join(absDir, path)

			// 检查文件是否存在
			info, err := os.Stat(fullPath)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("File not found: %s\n", fullPath)
					http.Error(w, "File not found", http.StatusNotFound)
				} else {
					fmt.Printf("Error accessing file: %v\n", err)
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}
				return
			}

			// 检查是否是文件（不是目录）
			if info.IsDir() {
				http.Error(w, "Path is directory", http.StatusNotFound)
				return
			}

			// 根据文件扩展名设置Content-Type
			ext := filepath.Ext(path)
			switch ext {
			case ".png":
				w.Header().Set("Content-Type", "image/png")
			case ".jpg", ".jpeg":
				w.Header().Set("Content-Type", "image/jpeg")
			case ".gif":
				w.Header().Set("Content-Type", "image/gif")
			case ".css":
				w.Header().Set("Content-Type", "text/css")
			case ".js":
				w.Header().Set("Content-Type", "application/javascript")
			}

			// 使用http.ServeFile提供文件服务
			http.ServeFile(w, r, fullPath)
		})

		// 使用通配符匹配所有静态文件请求
		server.AddRoute(rest.Route{
			Method:  http.MethodGet,
			Path:    "/common/static/*",
			Handler: staticHandler,
		})
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
