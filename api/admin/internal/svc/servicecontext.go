package svc

import (
	"admin/internal/config"
	"common/common"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	GreetClient common.CommonServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: c.Etcd,
	})
	return &ServiceContext{
		Config:      c,
		GreetClient: common.NewCommonServiceClient(conn.Conn()),
	}
}
