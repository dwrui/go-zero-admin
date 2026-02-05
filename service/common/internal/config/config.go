package config

import "github.com/zeromicro/go-zero/zrpc"

type App struct {
	LoginCaptcha bool `json:"loginCaptcha"`
}
type Config struct {
	zrpc.RpcServerConf
	App   App `json:"App"` // 或者直接嵌套
	Mysql struct {
		Host        string `json:"host"`
		Port        int    `json:"port"`
		Database    string `json:"database"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Charset     string `json:"charset"`
		TablePrefix string `json:"tablePrefix"`
	} `json:"mysql"`
}
