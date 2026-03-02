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

type SaveRuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveRuleLogic {
	return &SaveRuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveRuleLogic) SaveRule(req *types.SaveRuleReq) (resp *types.SaveRuleResp, err error) {
	rpcResp, err := l.svcCtx.SystemRuleClient.Save(l.ctx, &system.SaveRuleRequest{
		Activemenu:         req.Activemenu,
		Component:          req.Component,
		Des:                req.Des,
		Hidechildreninmenu: req.Hidechildreninmenu,
		Hideinmenu:         req.Hideinmenu,
		Icon:               req.Icon,
		Id:                 req.Id,
		Isext:              req.Isext,
		Keepalive:          req.Keepalive,
		Locale:             req.Locale,
		Noaffix:            req.Noaffix,
		Onlypage:           req.Onlypage,
		Path:               req.Path,
		Permission:         req.Permission,
		Pid:                req.Pid,
		Redirect:           req.Redirect,
		Requiresauth:       req.Requiresauth,
		Routename:          req.Routename,
		Routepath:          req.Routepath,
		Title:              req.Title,
		Type:               req.Ruletype,
		Weigh:              req.Weigh,
	})
	if err != nil {
		return nil, err
	}

	return &types.SaveRuleResp{
		Id: rpcResp.Id,
	}, nil
}
