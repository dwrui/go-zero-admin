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

type SaveConfigValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveConfigValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveConfigValueLogic {
	return &SaveConfigValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveConfigValueLogic) SaveConfigValue(req *types.SaveConfigValueReq) error {
	_, err := l.svcCtx.ConfigItemClient.SaveValue(l.ctx, &configcenter.SaveConfigValueRequest{
		CategoryKey:  req.CategoryKey,
		ConfigValues: req.ConfigValues,
	})
	if err != nil {
		l.Errorf("保存配置值失败: %v", err)
		return err
	}

	return nil
}
