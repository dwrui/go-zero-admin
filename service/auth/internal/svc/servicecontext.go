package svc

import (
	"auth/internal/config"
	"github.com/dwrui/go-zero-admin/pkg/utils/db"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	DB     *db.DBManager
}

func NewServiceContext(c config.Config) *ServiceContext {
	//初始化redis
	redisConf := redis.RedisConf{
		Host: c.RedisConf.Host,
		Type: c.RedisConf.Type,
		Pass: c.RedisConf.Pass,
		Tls:  c.RedisConf.Tls,
	}
	rds, _ := redis.NewRedis(redisConf)
	// 初始化数据库工具
	dbConfig := db.DBConfig{
		Host:        c.Mysql.Host,
		Port:        c.Mysql.Port,
		Database:    c.Mysql.Database,
		Username:    c.Mysql.Username,
		Password:    c.Mysql.Password,
		Charset:     "utf8mb4",
		TablePrefix: c.Mysql.TablePrefix,
	}
	db.InitDB(dbConfig)
	// 初始化数据库管理器
	return &ServiceContext{
		Config: c,
		Redis:  rds,
		DB:     db.GetDB(),
	}
}
