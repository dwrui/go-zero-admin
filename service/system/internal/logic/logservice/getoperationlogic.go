package logservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOperationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationLogic {
	return &GetOperationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOperationLogic) GetOperation(in *system.GetOperationRequest) (*system.GetOperationResponse, error) {
	// todo: add your logic here and delete this line

	return &system.GetOperationResponse{}, nil
}
