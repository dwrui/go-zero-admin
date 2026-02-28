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

type GetConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigListLogic {
	return &GetConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigListLogic) GetConfigList(req *types.GetConfigListReq) (resp *types.GetConfigListResp, err error) {
	rpcResp, err := l.svcCtx.ConfigItemClient.GetList(l.ctx, &configcenter.GetConfigListRequest{
		CategoryKey: req.CategoryKey,
		ConfigKey:   req.ConfigKey,
		ConfigName:  req.ConfigName,
		Status:      req.Status,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		l.Errorf("获取配置项列表失败: %v", err)
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

	return &types.GetConfigListResp{
		Items:    items,
		Total:    rpcResp.Total,
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
	}, nil
}
