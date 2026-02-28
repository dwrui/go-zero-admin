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

type DeleteConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigLogic {
	return &DeleteConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigLogic) DeleteConfig(req *types.DeleteConfigReq) error {
	_, err := l.svcCtx.ConfigItemClient.Delete(l.ctx, &configcenter.DeleteConfigRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("删除配置项失败: %v", err)
		return err
	}

	return nil
}
