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

type GetCategoryAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryAllLogic {
	return &GetCategoryAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryAllLogic) GetCategoryAll() (resp []types.ConfigCategoryData, err error) {
	rpcResp, err := l.svcCtx.ConfigCenterClient.GetAll(l.ctx, &configcenter.GetCategoryAllRequest{})
	if err != nil {
		l.Errorf("获取所有配置分类失败: %v", err)
		return nil, err
	}

	items := make([]types.ConfigCategoryData, 0)
	for _, item := range rpcResp.Items {
		items = append(items, types.ConfigCategoryData{
			Id:           item.Id,
			CategoryKey:  item.CategoryKey,
			CategoryName: item.CategoryName,
			Description:  item.Description,
			SortOrder:    item.SortOrder,
			IsSystem:     item.IsSystem,
			CreateTime:   item.CreateTime,
			UpdateTime:   item.UpdateTime,
		})
	}

	return items, nil
}
