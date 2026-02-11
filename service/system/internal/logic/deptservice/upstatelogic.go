package deptservicelogic

import (
	"context"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStateLogic {
	return &UpStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpStateLogic) UpState(in *system.UpStatusDeptRequest) (*system.UpStatusDeptResponse, error) {
	_, err := model.UpStatusDept(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	return &system.UpStatusDeptResponse{}, nil
}
