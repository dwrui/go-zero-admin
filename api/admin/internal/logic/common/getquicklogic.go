// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package common

import (
	"context"

	"admin/internal/svc"
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
	return nil, nil
}
