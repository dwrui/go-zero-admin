package authservicelogic

import (
	"context"
	"errors"
	"github.com/dwrui/go-zero-admin/pkg/utils/jwt"

	"auth/auth"
	"auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(in *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	// todo:
	// 校验 RefreshToken 的有效性
	jwtConfig := jwt.JwtConfig{
		AccessSecret: l.svcCtx.Config.Jwt.AccessSecret,
		AccessExpire: l.svcCtx.Config.Jwt.AccessExpire,
	}
	claims, err := jwt.ParseToken(jwtConfig, in.Token)
	if err != nil {
		return nil, errors.New("无效的 RefreshToken")
	}

	// 生成新的 AccessToken
	newToken, err := jwt.GenerateToken(jwtConfig, claims.Data)
	if err != nil {
		return nil, errors.New("生成新 Token 失败")
	}
	return &auth.RefreshTokenResponse{
		NewToken: newToken,
	}, nil
}
