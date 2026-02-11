package ruleservicelogic

import (
	"context"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStatusLogic {
	return &UpStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpStatusLogic) UpStatus(in *system.UpStatusRuleRequest) (*system.UpStatusRuleResponse, error) {
	err := model.UpStatus(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	return &system.UpStatusRuleResponse{}, nil
}
