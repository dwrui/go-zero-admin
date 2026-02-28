package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *configcenter.UpdateConfigRequest) (*configcenter.UpdateConfigResponse, error) {
	// 检查配置项是否存在
	item, err := model.GetConfigDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("配置项不存在")
	}

	// 如果修改了分类或配置键，检查是否冲突
	if item.CategoryKey != in.CategoryKey || item.ConfigKey != in.ConfigKey {
		existItem, err := model.GetConfigByKey(l.ctx, l.svcCtx, in.CategoryKey, in.ConfigKey)
		if err != nil {
			return nil, err
		}
		if existItem != nil && existItem.Id != int64(in.Id) {
			return nil, errors.New("该分类下配置键已存在")
		}
	}

	// 检查新分类是否存在
	if item.CategoryKey != in.CategoryKey {
		category, err := model.GetCategoryByKey(l.ctx, l.svcCtx, in.CategoryKey)
		if err != nil {
			return nil, err
		}
		if category == nil {
			return nil, errors.New("分类不存在")
		}
	}

	err = model.UpdateConfig(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}

	return &configcenter.UpdateConfigResponse{}, nil
}
