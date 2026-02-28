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

type GetConfigDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigDetailLogic {
	return &GetConfigDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigDetailLogic) GetConfigDetail(req *types.GetConfigDetailReq) (resp *types.ConfigItemData, err error) {
	rpcResp, err := l.svcCtx.ConfigItemClient.GetDetail(l.ctx, &configcenter.GetConfigDetailRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("获取配置项详情失败: %v", err)
		return nil, err
	}

	options := make([]types.ConfigOptionItem, 0)
	for _, opt := range rpcResp.Data.Options {
		options = append(options, types.ConfigOptionItem{
			Label: opt.Label,
			Value: opt.Value,
		})
	}

	return &types.ConfigItemData{
		Id:             rpcResp.Data.Id,
		CategoryKey:    rpcResp.Data.CategoryKey,
		ConfigKey:      rpcResp.Data.ConfigKey,
		ConfigName:     rpcResp.Data.ConfigName,
		ConfigType:     rpcResp.Data.ConfigType,
		ConfigValue:    rpcResp.Data.ConfigValue,
		DefaultValue:   rpcResp.Data.DefaultValue,
		Description:    rpcResp.Data.Description,
		Options:        options,
		ValidationRule: rpcResp.Data.ValidationRule,
		Placeholder:    rpcResp.Data.Placeholder,
		IsRequired:     rpcResp.Data.IsRequired,
		IsSecret:       rpcResp.Data.IsSecret,
		SortOrder:      rpcResp.Data.SortOrder,
		Status:         rpcResp.Data.Status,
		CreatedTime:    rpcResp.Data.CreatedTime,
		UpdatedTime:    rpcResp.Data.UpdatedTime,
	}, nil
}
