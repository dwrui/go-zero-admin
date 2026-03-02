// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package rule

import (
	"context"
	"encoding/json"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRuleParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRuleParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRuleParentLogic {
	return &GetRuleParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRuleParentLogic) GetRuleParent(req *types.GetRuleParentReq) (any, error) {
	rpcResp, err := l.svcCtx.SystemRuleClient.GetParent(l.ctx, &system.GetRuleParentRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var respMap map[string]interface{}

	err = json.Unmarshal([]byte(rpcResp.Data), &respMap)
	if err != nil {
		return nil, err
	}
	return respMap, nil
}
