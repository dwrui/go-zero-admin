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

type UpdateConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigLogic {
	return &UpdateConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigLogic) UpdateConfig(req *types.UpdateConfigReq) error {
	_, err := l.svcCtx.ConfigItemClient.Update(l.ctx, &configcenter.UpdateConfigRequest{
		Id:             req.Id,
		CategoryKey:    req.CategoryKey,
		ConfigKey:      req.ConfigKey,
		ConfigName:     req.ConfigName,
		ConfigType:     req.ConfigType,
		ConfigValue:    req.ConfigValue,
		DefaultValue:   req.DefaultValue,
		Description:    req.Description,
		Options:        req.Options,
		ValidationRule: req.ValidationRule,
		Placeholder:    req.Placeholder,
		IsRequired:     req.IsRequired,
		IsSecret:       req.IsSecret,
		SortOrder:      req.SortOrder,
		Status:         req.Status,
	})
	if err != nil {
		l.Errorf("更新配置项失败: %v", err)
		return err
	}

	return nil
}
