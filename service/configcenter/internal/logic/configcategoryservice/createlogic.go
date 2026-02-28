package configcategoryservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *configcenter.CreateCategoryRequest) (*configcenter.CreateCategoryResponse, error) {
	// 检查分类标识是否已存在
	existCategory, err := model.GetCategoryByKey(l.ctx, l.svcCtx, in.CategoryKey)
	if err != nil {
		return nil, err
	}
	if existCategory != nil {
		return nil, errors.New("分类标识已存在")
	}

	id, err := model.CreateCategory(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}

	return &configcenter.CreateCategoryResponse{
		Id: ga.Uint64(id),
	}, nil
}
