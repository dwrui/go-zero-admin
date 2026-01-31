package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Host        string `json:"host"`
		Port        int    `json:"port"`
		Database    string `json:"database"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Charset     string `json:"charset"`
		TablePrefix string `json:"tablePrefix"`
	} `json:"mysql"`
	Jwt struct {
		AccessSecret string `json:"accessSecret"`
		AccessExpire int64  `json:"accessExpire"`
	} `json:"jwt"`
}
