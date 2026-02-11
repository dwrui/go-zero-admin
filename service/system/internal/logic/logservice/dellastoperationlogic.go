package logservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLastOperationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLastOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLastOperationLogic {
	return &DelLastOperationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLastOperationLogic) DelLastOperation(in *system.DelLastOperationRequest) (*system.DelLastOperationResponse, error) {
	// todo: add your logic here and delete this line

	return &system.DelLastOperationResponse{}, nil
}
