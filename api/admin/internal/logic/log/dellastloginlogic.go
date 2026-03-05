// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package log

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLastLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLastLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLastLoginLogic {
	return &DelLastLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLastLoginLogic) DelLastLogin(req *types.DelLastLoginReq) (resp *types.DelLastLoginResp, err error) {
	_, err = l.svcCtx.SystemLogClient.DelLastLogin(l.ctx, &system.DelLastLoginRequest{})
	if err != nil {
		return nil, err
	}

	return &types.DelLastLoginResp{}, nil
}
