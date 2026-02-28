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

type UpdateConfigStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigStatusLogic {
	return &UpdateConfigStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigStatusLogic) UpdateConfigStatus(req *types.UpdateConfigStatusReq) error {
	_, err := l.svcCtx.ConfigItemClient.UpdateStatus(l.ctx, &configcenter.UpdateConfigStatusRequest{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		l.Errorf("更新配置项状态失败: %v", err)
		return err
	}

	return nil
}
