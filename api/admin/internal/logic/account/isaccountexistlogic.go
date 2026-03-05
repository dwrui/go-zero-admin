package account

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsAccountExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsAccountExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsAccountExistLogic {
	return &IsAccountExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsAccountExistLogic) IsAccountExist(req *types.IsAccountExistReq) (resp *types.IsAccountExistResp, err error) {
	_, err = l.svcCtx.SystemAccountClient.Isaccountexist(l.ctx, &system.IsAccountExistRequest{
		Id:       req.Id,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	return &types.IsAccountExistResp{}, nil
}