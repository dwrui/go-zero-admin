package user

import (
	"admin/internal/svc"
	_ "admin/internal/types"
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
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

func (l *GetUserInfoLogic) GetUserInfo() (any, error) {
	userId := ga.Uint64(l.ctx.Value("id"))
	resp, err := l.svcCtx.UserClient.GetUserinfo(l.ctx, &user.GetUserinfoRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
