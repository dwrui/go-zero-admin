package config

import "github.com/zeromicro/go-zero/zrpc"

type App struct {
	LoginCaptcha bool `json:"loginCaptcha"`
}
type Config struct {
	zrpc.RpcServerConf
	App App `json:"App"` // 或者直接嵌套
}
