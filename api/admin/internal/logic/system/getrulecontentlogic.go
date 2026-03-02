// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRuleContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRuleContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRuleContentLogic {
	return &GetRuleContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRuleContentLogic) GetRuleContent(req *types.GetRuleContentReq) (resp *types.GetRuleContentResp, err error) {
	rpcResp, err := l.svcCtx.SystemRuleClient.GetContent(l.ctx, &system.GetRuleContentRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetRuleContentResp{
		Activemenu:         rpcResp.Activemenu,
		Component:          rpcResp.Component,
		Des:                rpcResp.Des,
		Hidechildreninmenu: rpcResp.Hidechildreninmenu,
		Hideinmenu:         rpcResp.Hideinmenu,
		Icon:               rpcResp.Icon,
		Id:                 rpcResp.Id,
		Isext:              rpcResp.Isext,
		Keepalive:          rpcResp.Keepalive,
		Locale:             rpcResp.Locale,
		Noaffix:            rpcResp.Noaffix,
		Onlypage:           rpcResp.Onlypage,
		Path:               rpcResp.Path,
		Permission:         rpcResp.Permission,
		Pid:                rpcResp.Pid,
		Redirect:           rpcResp.Redirect,
		Requiresauth:       rpcResp.Requiresauth,
		Routename:          rpcResp.Routename,
		Routepath:          rpcResp.Routepath,
		Title:              rpcResp.Title,
		Ruletype:           rpcResp.Type,
		Weigh:              rpcResp.Weigh,
		CreateTime:         rpcResp.CreateTime,
		Status:             rpcResp.Status,
		Uid:                rpcResp.Uid,
	}, nil
}
