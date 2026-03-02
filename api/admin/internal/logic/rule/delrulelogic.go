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

type DelRuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRuleLogic {
	return &DelRuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelRuleLogic) DelRule(req *types.DelRuleReq) (resp *types.DelRuleResp, err error) {
	_, err = l.svcCtx.SystemRuleClient.Del(l.ctx, &system.DelRuleRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DelRuleResp{}, nil
}
