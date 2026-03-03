package roleservicelogic

import (
	"context"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveLogic) Save(in *system.SaveRoleRequest) (*system.SaveRoleResponse, error) {
	// todo: add your logic here and delete this line
	roleId, err := model.SaveRole(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	return &system.SaveRoleResponse{
		Id: roleId,
	}, nil
}
