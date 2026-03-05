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

type DelLastOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLastOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLastOperationLogic {
	return &DelLastOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLastOperationLogic) DelLastOperation(req *types.DelLastOperationReq) (resp *types.DelLastOperationResp, err error) {
	_, err = l.svcCtx.SystemLogClient.DelLastOperation(l.ctx, &system.DelLastOperationRequest{})
	if err != nil {
		return nil, err
	}

	return &types.DelLastOperationResp{}, nil
}
