package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Host        string
		Port        int
		Database    string
		Username    string
		Password    string
		Charset     string
		TablePrefix string
	} `json:"mysql"`
}
