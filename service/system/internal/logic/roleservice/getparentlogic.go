package roleservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParentLogic {
	return &GetParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetParentLogic) GetParent(in *system.GetRoleParentRequest) (*system.GetRoleParentResponse, error) {
	// todo: add your logic here and delete this line

	return &system.GetRoleParentResponse{}, nil
}
