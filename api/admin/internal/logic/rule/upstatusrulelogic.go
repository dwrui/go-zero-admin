// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package rule

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStatusRuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpStatusRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStatusRuleLogic {
	return &UpStatusRuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpStatusRuleLogic) UpStatusRule(req *types.UpStatusRuleReq) (resp *types.UpStatusRuleResp, err error) {
	_, err = l.svcCtx.SystemRuleClient.UpStatus(l.ctx, &system.UpStatusRuleRequest{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpStatusRuleResp{}, nil
}
