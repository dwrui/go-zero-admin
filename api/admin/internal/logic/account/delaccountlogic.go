package account

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAccountLogic {
	return &DelAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAccountLogic) DelAccount(req *types.DelAccountReq) (resp *types.DelAccountResp, err error) {
	_, err = l.svcCtx.SystemAccountClient.Del(l.ctx, &system.DelAccountRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DelAccountResp{}, nil
}