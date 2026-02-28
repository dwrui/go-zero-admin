// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"admin/grpc-client/configcenter"
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryReq) error {
	_, err := l.svcCtx.ConfigCenterClient.Create(l.ctx, &configcenter.CreateCategoryRequest{
		CategoryKey:  req.CategoryKey,
		CategoryName: req.CategoryName,
		Description:  req.Description,
		SortOrder:    req.SortOrder,
	})
	if err != nil {
		l.Errorf("创建配置分类失败: %v", err)
		return err
	}

	return nil
}
