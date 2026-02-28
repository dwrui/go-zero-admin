package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetValueLogic {
	return &GetValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetValueLogic) GetValue(in *configcenter.GetConfigValueRequest) (*configcenter.GetConfigValueResponse, error) {
	// 检查分类是否存在
	category, err := model.GetCategoryByKey(l.ctx, l.svcCtx, in.CategoryKey)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 获取配置值
	values, err := model.GetConfigValue(l.ctx, l.svcCtx, in.CategoryKey)
	if err != nil {
		return nil, err
	}

	var items []*configcenter.ConfigValueItem
	for key, value := range values {
		items = append(items, &configcenter.ConfigValueItem{
			ConfigKey:   key,
			ConfigValue: value,
		})
	}

	return &configcenter.GetConfigValueResponse{
		Items: items,
	}, nil
}
