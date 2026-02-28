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

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) error {
	_, err := l.svcCtx.ConfigCenterClient.Update(l.ctx, &configcenter.UpdateCategoryRequest{
		Id:           req.Id,
		CategoryName: req.CategoryName,
		Description:  req.Description,
		SortOrder:    req.SortOrder,
	})
	if err != nil {
		l.Errorf("更新配置分类失败: %v", err)
		return err
	}

	return nil
}
