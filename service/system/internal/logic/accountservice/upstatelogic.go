package accountservicelogic

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

func (l *UpStateLogic) UpState(in *system.UpStatusAccountRequest) (*system.UpStatusAccountResponse, error) {
	err := model.UpStatusAccount(l.ctx, l.svcCtx, in.Id, in.Status)
	if err != nil {
		return nil, err
	}
	return &system.UpStatusAccountResponse{}, nil
}
