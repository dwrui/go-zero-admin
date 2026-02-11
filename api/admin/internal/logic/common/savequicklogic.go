// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package common

import (
	"admin/grpc-client/common"
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveQuickLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveQuickLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveQuickLogic {
	return &SaveQuickLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveQuickLogic) SaveQuick(req *types.SaveQuickReq) (any, error) {
	resp, err := l.svcCtx.DashboardClient.SaveQuick(l.ctx, &common.SaveQuickRequest{
		Icon:       req.Icon,
		Id:         req.Id,
		Name:       req.Name,
		PathUrl:    req.PathUrl,
		Type:       req.ReqType,
		Weigh:      req.Weigh,
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
		Uid:        ga.Uint64(l.ctx.Value("user_id")),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
