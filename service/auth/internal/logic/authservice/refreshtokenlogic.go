package authservicelogic

import (
	"auth/internal/model"
	"context"
	"errors"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
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
	userId := in.UserId
	userInfo, err := model.GetUserInfo(l.ctx, l.svcCtx, userId, "id,account_id,business_id")
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	jwtConfig := jwt.JwtConfig{
		AccessSecret: l.svcCtx.Config.Jwt.AccessSecret,
		AccessExpire: l.svcCtx.Config.Jwt.AccessExpire,
	}
	token, err := jwt.GenerateToken(jwtConfig, ga.Map{"id": userInfo.Id, "account_id": userInfo.AccountId, "business_id": userInfo.BusinessId})
	if err != nil {
		return nil, errors.New("生成新 Token 失败")
	}
	return &auth.RefreshTokenResponse{
		NewToken: token,
	}, nil
}
