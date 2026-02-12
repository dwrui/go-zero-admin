// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"admin/internal/types"
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"system/system"

	"admin/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogic {
	return &GetLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginLogic) GetLogin(req *types.GetLoginReq) (any, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.SystemLogClient.GetLogin(l.ctx, &system.GetLogListRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Status:     req.Status,
		User:       req.User,
		CreateTime: req.Create_time,
		Ip:         req.Ip,
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
