// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package dept

import (
	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"
	"context"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveDeptLogic {
	return &SaveDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveDeptLogic) SaveDept(req *types.SaveDeptReq) (resp *types.SaveDeptResp, err error) {
	rpcReq := &system.SaveDeptRequest{
		Id:         req.Id,
		AccountId:  ga.Uint64(l.ctx.Value("user_id")),
		Name:       req.Name,
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
		Pid:        req.Pid,
		Remark:     req.Remark,
		Weigh:      req.Weigh,
	}
	rpcResp, err := l.svcCtx.SystemDeptClient.Save(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return &types.SaveDeptResp{
		Id: rpcResp.Id,
	}, nil
}
