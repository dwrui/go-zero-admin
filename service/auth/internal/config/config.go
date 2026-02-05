package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Jwt struct {
		AccessSecret string
		AccessExpire int64
	} `json:"jwt"`
	Redis struct {
		Host string
		Type string
		Pass string
		Tls  bool
	} `json:"redis"`
}
