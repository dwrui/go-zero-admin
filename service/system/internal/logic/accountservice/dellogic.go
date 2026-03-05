package accountservicelogic

import (
	"context"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogic) Del(in *system.DelAccountRequest) (*system.DelAccountResponse, error) {
	err := model.DelAccount(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	return &system.DelAccountResponse{}, nil
}
