package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *configcenter.DeleteConfigRequest) (*configcenter.DeleteConfigResponse, error) {
	// 检查配置项是否存在
	item, err := model.GetConfigDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("配置项不存在")
	}

	err = model.DeleteConfig(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}

	return &configcenter.DeleteConfigResponse{}, nil
}
