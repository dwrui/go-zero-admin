package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByCategoryLogic {
	return &GetByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByCategoryLogic) GetByCategory(in *configcenter.GetConfigByCategoryRequest) (*configcenter.GetConfigByCategoryResponse, error) {
	// 获取分类信息
	category, err := model.GetCategoryByKey(l.ctx, l.svcCtx, in.CategoryKey)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 获取配置项列表
	list, err := model.GetConfigByCategory(l.ctx, l.svcCtx, in.CategoryKey)
	if err != nil {
		return nil, err
	}

	var items []*configcenter.ConfigItemData
	for _, item := range list {
		items = append(items, convertToConfigItemData(item))
	}

	return &configcenter.GetConfigByCategoryResponse{
		CategoryName:        category.CategoryName,
		CategoryDescription: category.Description.String,
		Items:               items,
	}, nil
}
