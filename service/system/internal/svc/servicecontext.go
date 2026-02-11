package svc

import (
	"github.com/dwrui/go-zero-admin/pkg/utils/db"
	"system/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *db.DBManager
}

func NewServiceContext(c config.Config) *ServiceContext {
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
	// 创建数据库连接
	db.InitDB(dbConfig)
	return &ServiceContext{
		Config: c,
		DB:     db.GetDB(),
	}
}
