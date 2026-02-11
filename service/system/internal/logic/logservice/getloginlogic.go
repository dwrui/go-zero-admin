package logservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogic {
	return &GetLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLoginLogic) GetLogin(in *system.GetLogListRequest) (*system.GetLogListResponse, error) {
	// todo: add your logic here and delete this line

	return &system.GetLogListResponse{}, nil
}
