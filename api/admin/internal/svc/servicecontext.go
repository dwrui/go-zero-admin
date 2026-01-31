package svc

import (
	"admin/grpc-client/common"
	"admin/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"user/user"
)

type ServiceContext struct {
	Config       config.Config
	CommonClient common.CommonServiceClient
	UserClient   user.UserServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	//common client链接
	commonConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: c.CommonEtcd,
	})
	////user client链接
	userConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: c.UserEtcd,
	})
	return &ServiceContext{
		Config:       c,
		CommonClient: common.NewCommonServiceClient(commonConn.Conn()),
		UserClient:   user.NewUserServiceClient(userConn.Conn()),
	}
}
