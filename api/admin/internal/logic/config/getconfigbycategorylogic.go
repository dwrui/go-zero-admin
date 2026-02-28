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

type GetConfigByCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigByCategoryLogic {
	return &GetConfigByCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigByCategoryLogic) GetConfigByCategory(req *types.GetConfigByCategoryReq) (resp map[string]interface{}, err error) {
	rpcResp, err := l.svcCtx.ConfigItemClient.GetByCategory(l.ctx, &configcenter.GetConfigByCategoryRequest{
		CategoryKey: req.CategoryKey,
	})
	if err != nil {
		l.Errorf("根据分类获取配置项失败: %v", err)
		return nil, err
	}

	items := make([]types.ConfigItemData, 0)
	for _, item := range rpcResp.Items {
		options := make([]types.ConfigOptionItem, 0)
		for _, opt := range item.Options {
			options = append(options, types.ConfigOptionItem{
				Label: opt.Label,
				Value: opt.Value,
			})
		}

		items = append(items, types.ConfigItemData{
			Id:             item.Id,
			CategoryKey:    item.CategoryKey,
			ConfigKey:      item.ConfigKey,
			ConfigName:     item.ConfigName,
			ConfigType:     item.ConfigType,
			ConfigValue:    item.ConfigValue,
			DefaultValue:   item.DefaultValue,
			Description:    item.Description,
			Options:        options,
			ValidationRule: item.ValidationRule,
			Placeholder:    item.Placeholder,
			IsRequired:     item.IsRequired,
			IsSecret:       item.IsSecret,
			SortOrder:      item.SortOrder,
			Status:         item.Status,
			CreatedTime:    item.CreatedTime,
			UpdatedTime:    item.UpdatedTime,
		})
	}

	return map[string]interface{}{
		"category_name":        rpcResp.CategoryName,
		"category_description": rpcResp.CategoryDescription,
		"items":                items,
	}, nil
}
