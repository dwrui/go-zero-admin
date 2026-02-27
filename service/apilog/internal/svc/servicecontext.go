package svc

import (
	"apilog/internal/config"

	"github.com/dwrui/go-zero-admin/pkg/utils/db"
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
	db.InitDB(dbConfig)
	return &ServiceContext{
		Config: c,
		DB:     db.GetDB(),
	}
}
