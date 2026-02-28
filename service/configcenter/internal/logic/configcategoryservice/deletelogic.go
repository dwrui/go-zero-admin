package configcategoryservicelogic

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

func (l *DeleteLogic) Delete(in *configcenter.DeleteCategoryRequest) (*configcenter.DeleteCategoryResponse, error) {
	// 检查分类是否存在
	category, err := model.GetCategoryDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 系统内置分类不允许删除
	if category.IsSystem == 1 {
		return nil, errors.New("系统内置分类不允许删除")
	}

	// 检查分类下是否有配置项
	items, err := model.GetConfigByCategory(l.ctx, l.svcCtx, category.CategoryKey)
	if err != nil {
		return nil, err
	}
	if len(items) > 0 {
		return nil, errors.New("该分类下存在配置项，请先删除配置项")
	}

	err = model.DeleteCategory(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}

	return &configcenter.DeleteCategoryResponse{}, nil
}
