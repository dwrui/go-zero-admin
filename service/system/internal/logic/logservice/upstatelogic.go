package logservicelogic

import (
	"context"

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

func (l *UpStateLogic) UpState(in *system.UpStatusLogRequest) (*system.UpStatusLogResponse, error) {
	// todo: add your logic here and delete this line

	return &system.UpStatusLogResponse{}, nil
}
