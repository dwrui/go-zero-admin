package ruleservicelogic

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoutesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoutesLogic {
	return &GetRoutesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoutesLogic) GetRoutes(in *system.GetRoutesRequest) (*system.GetRoutesResponse, error) {
	//获取所有的路由信息
	rule, err := model.GetRoutesAll(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	ruleJson, _ := json.Marshal(rule)
	return &system.GetRoutesResponse{
		Data: ga.String(ruleJson),
	}, nil
}
