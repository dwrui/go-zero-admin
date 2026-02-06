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
	RedisConf struct {
		Host string
		Type string
		Pass string
		Tls  bool
	} `json:"redisConf"`
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
