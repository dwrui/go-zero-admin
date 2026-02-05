package svc

import (
	"auth/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化redis
	redisConf := redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
		Tls:  c.Redis.Tls,
	}
	rds, _ := redis.NewRedis(redisConf)
	return &ServiceContext{
		Config: c,
		Redis:  rds,
	}
}
