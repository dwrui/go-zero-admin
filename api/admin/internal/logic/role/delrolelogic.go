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

type DelRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRoleLogic {
	return &DelRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelRoleLogic) DelRole(req *types.DelRoleReq) (resp *types.DelRoleResp, err error) {
	rpcReq := &system.DelRoleRequest{
		Id: req.Id,
	}
	_, err = l.svcCtx.SystemRoleClient.Del(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return &types.DelRoleResp{}, nil
}
