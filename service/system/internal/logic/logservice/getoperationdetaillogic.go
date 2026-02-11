package logservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOperationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOperationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationDetailLogic {
	return &GetOperationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOperationDetailLogic) GetOperationDetail(in *system.GetOperationDetailRequest) (*system.GetOperationDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &system.GetOperationDetailResponse{}, nil
}
