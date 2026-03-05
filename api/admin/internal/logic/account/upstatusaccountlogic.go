package account

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStatusAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpStatusAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStatusAccountLogic {
	return &UpStatusAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpStatusAccountLogic) UpStatusAccount(req *types.UpStatusAccountReq) (resp *types.UpStatusAccountResp, err error) {
	_, err = l.svcCtx.SystemAccountClient.UpState(l.ctx, &system.UpStatusAccountRequest{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpStatusAccountResp{}, nil
}