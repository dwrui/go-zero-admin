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

type GetConfigValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigValueLogic {
	return &GetConfigValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigValueLogic) GetConfigValue(req *types.GetConfigValueReq) (resp map[string]string, err error) {
	rpcResp, err := l.svcCtx.ConfigItemClient.GetValue(l.ctx, &configcenter.GetConfigValueRequest{
		CategoryKey: req.CategoryKey,
	})
	if err != nil {
		l.Errorf("获取配置值失败: %v", err)
		return nil, err
	}

	result := make(map[string]string)
	for _, item := range rpcResp.Items {
		result[item.ConfigKey] = item.ConfigValue
	}

	return result, nil
}
