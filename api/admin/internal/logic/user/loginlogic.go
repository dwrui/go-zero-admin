package user

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"user/user"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/**
 * 登录
 * @param req
 * @return any
 * @return error
 */
func (l *LoginLogic) Login(req *types.LoginReq, reqs ga.Map) (any, error) {
	// 调用user.rpc服务的Login方法
	rpcResp, err := l.svcCtx.UserClient.Login(l.ctx, &user.LoginRequest{
		Username:  req.Username,
		Password:  req.Password,
		Codeid:    req.Codeid,
		Captcha:   req.Captcha,
		ClientIp:  ga.String(reqs["ip"]),
		UserAgent: ga.String(reqs["user_agent"]),
	})
	if err != nil {
		return nil, err
	}
	return rpcResp, nil
}
