package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStatusLogic) UpdateStatus(in *configcenter.UpdateConfigStatusRequest) (*configcenter.UpdateConfigStatusResponse, error) {
	// 检查配置项是否存在
	item, err := model.GetConfigDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("配置项不存在")
	}

	err = model.UpdateConfigStatus(l.ctx, l.svcCtx, in.Id, in.Status)
	if err != nil {
		return nil, err
	}

	return &configcenter.UpdateConfigStatusResponse{}, nil
}
