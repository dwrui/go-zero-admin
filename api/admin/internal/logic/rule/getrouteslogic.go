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

type GetRoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoutesLogic {
	return &GetRoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoutesLogic) GetRoutes(req *types.GetRoutesReq) (any, error) {
	rpcResp, err := l.svcCtx.SystemRuleClient.GetRoutes(l.ctx, &system.GetRoutesRequest{})
	if err != nil {
		return nil, err
	}
	var respMap []map[string]interface{}
	_ = json.Unmarshal([]byte(rpcResp.Data), &respMap)
	return respMap, nil
}
