package roleservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStatusLogic {
	return &UpStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpStatusLogic) UpStatus(in *system.UpStatusRoleRequest) (*system.UpStatusRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &system.UpStatusRoleResponse{}, nil
}
