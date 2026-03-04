// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package role

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStatusRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpStatusRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStatusRoleLogic {
	return &UpStatusRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpStatusRoleLogic) UpStatusRole(req *types.UpStatusRoleReq) (resp *types.UpStatusRoleResp, err error) {
	rpcReq := &system.UpStatusRoleRequest{
		Id:     req.Id,
		Status: req.Status,
	}
	_, err = l.svcCtx.SystemRoleClient.UpStatus(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return &types.UpStatusRoleResp{}, nil
}
