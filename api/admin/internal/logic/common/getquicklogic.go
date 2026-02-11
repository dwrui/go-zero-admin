// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package common

import (
	"admin/grpc-client/common"
	"admin/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuickLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuickLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuickLogic {
	return &GetQuickLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuickLogic) GetQuick() (any, error) {
	resp, err := l.svcCtx.DashboardClient.GetQuick(l.ctx, &common.GetQuickRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
