package user

import (
	"admin/internal/svc"
	_ "admin/internal/types"
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/jwt"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(token string) (any, error) {
	jwtConfig := jwt.JwtConfig{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
	}
	claims, err := jwt.ParseToken(jwtConfig, token)
	if err != nil {
		return nil, err
	}
	resp, err := l.svcCtx.UserClient.GetUserinfo(l.ctx, &user.GetUserinfoRequest{
		UserId: ga.Uint64(claims.Data["id"]),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
