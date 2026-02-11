package svc

import (
	"admin/grpc-client/common"
	"admin/internal/config"
	"auth/auth"
	"context"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"user/user"
)

type ServiceContext struct {
	Config          config.Config
	CommonClient    common.CommonServiceClient
	DashboardClient common.DashboardServiceClient
	UserClient      user.UserServiceClient
	AuthClient      auth.AuthServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	//common client链接
	commonConn := zrpc.MustNewClient(createRpcClientConf(c.CommonEtcd))
	//user client链接
	userConn := zrpc.MustNewClient(createRpcClientConf(c.UserEtcd))
	//auth client链接
	authConn := zrpc.MustNewClient(createRpcClientConf(c.AuthEtcd))
	return &ServiceContext{
		Config:          c,
		CommonClient:    common.NewCommonServiceClient(commonConn.Conn()),
		DashboardClient: common.NewDashboardServiceClient(commonConn.Conn()),
		UserClient:      user.NewUserServiceClient(userConn.Conn()),
		AuthClient:      auth.NewAuthServiceClient(authConn.Conn()),
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

// 统一鉴权方法
func (svcCtx *ServiceContext) CheckPermission(ctx context.Context, r *http.Request, token string, permission string) error {
	resp, err := svcCtx.AuthClient.CheckToken(ctx, &auth.CheckTokenRequest{
		Token:      token,
		Permission: permission,
	})
	if err != nil {
		return err
	}
	if resp.NewToken != "" {
		r.Header.Set("X-New-Token", resp.NewToken)
		return nil
	}
	return nil
}
