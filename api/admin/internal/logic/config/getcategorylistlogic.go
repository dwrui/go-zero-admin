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

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.GetCategoryListReq) (resp *types.GetCategoryListResp, err error) {
	rpcResp, err := l.svcCtx.ConfigCenterClient.GetList(l.ctx, &configcenter.GetCategoryListRequest{
		CategoryKey:  req.CategoryKey,
		CategoryName: req.CategoryName,
		Page:         req.Page,
		PageSize:     req.PageSize,
	})
	if err != nil {
		l.Errorf("获取配置分类列表失败: %v", err)
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

	return &types.GetCategoryListResp{
		Items:    items,
		Total:    rpcResp.Total,
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
	}, nil
}
