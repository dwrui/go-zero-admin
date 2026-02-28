package configcategoryservicelogic

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

func (l *UpdateLogic) Update(in *configcenter.UpdateCategoryRequest) (*configcenter.UpdateCategoryResponse, error) {
	// 检查分类是否存在
	category, err := model.GetCategoryDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 系统内置分类不允许修改
	if category.IsSystem == 1 {
		return nil, errors.New("系统内置分类不允许修改")
	}

	err = model.UpdateCategory(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}

	return &configcenter.UpdateCategoryResponse{}, nil
}
