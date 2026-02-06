package svc

import (
	"admin/grpc-client/common"
	"admin/internal/config"
	"auth/auth"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"user/user"
)

type ServiceContext struct {
	Config       config.Config
	CommonClient common.CommonServiceClient
	UserClient   user.UserServiceClient
	AuthClient   auth.AuthServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	//common client链接
	commonConn := zrpc.MustNewClient(createRpcClientConf(c.CommonEtcd))
	//user client链接
	userConn := zrpc.MustNewClient(createRpcClientConf(c.UserEtcd))
	//auth client链接
	authConn := zrpc.MustNewClient(createRpcClientConf(c.AuthEtcd))
	return &ServiceContext{
		Config:       c,
		CommonClient: common.NewCommonServiceClient(commonConn.Conn()),
		UserClient:   user.NewUserServiceClient(userConn.Conn()),
		AuthClient:   auth.NewAuthServiceClient(authConn.Conn()),
	}
}

// 统一RPC客户端配置
func createRpcClientConf(etcdConf discov.EtcdConf) zrpc.RpcClientConf {
	return zrpc.RpcClientConf{
		Etcd:     etcdConf,
		NonBlock: true, // 非阻塞模式
		Timeout:  3000, // 3秒超时
	}
}
