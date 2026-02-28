package configitemservicelogic

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

func (l *CreateLogic) Create(in *configcenter.CreateConfigRequest) (*configcenter.CreateConfigResponse, error) {
	// 检查分类是否存在
	category, err := model.GetCategoryByKey(l.ctx, l.svcCtx, in.CategoryKey)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 检查配置键是否已存在
	existItem, err := model.GetConfigByKey(l.ctx, l.svcCtx, in.CategoryKey, in.ConfigKey)
	if err != nil {
		return nil, err
	}
	if existItem != nil {
		return nil, errors.New("该分类下配置键已存在")
	}

	id, err := model.CreateConfig(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}

	return &configcenter.CreateConfigResponse{
		Id: ga.Uint64(id),
	}, nil
}
