package ruleservicelogic

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentLogic) GetContent(in *system.GetRuleContentRequest) (*system.GetRuleContentResponse, error) {
	rule, err := model.GetRuleOne(l.ctx, l.svcCtx, "activemenu,component,create_time,des,hidechildreninmenu,hideinmenu,icon,id,isext,keepalive,locale,noaffix,onlypage,path,permission,pid,redirect,requiresauth,routename,routepath,status,title,type,uid,weigh", in.Id)
	if err != nil {
		return nil, err
	}
	return &system.GetRuleContentResponse{
		Activemenu:         rule.Activemenu,
		Component:          rule.Component,
		CreateTime:         ga.String(rule.Createtime),
		Des:                rule.Des,
		Hidechildreninmenu: rule.Hidechildreninmenu,
		Hideinmenu:         rule.Hideinmenu,
		Icon:               rule.Icon,
		Id:                 rule.Id,
		Isext:              rule.Isext,
		Keepalive:          rule.Keepalive,
		Locale:             rule.Locale,
		Noaffix:            rule.Noaffix,
		Onlypage:           rule.Onlypage,
		Path:               rule.Path,
		Permission:         ga.String(rule.Permission),
		Pid:                rule.Pid,
		Redirect:           rule.Redirect,
		Requiresauth:       rule.Requiresauth,
		Routename:          rule.Routename,
		Routepath:          rule.Routepath,
		Status:             rule.Status,
		Title:              rule.Title,
		Type:               rule.Type,
		Uid:                rule.Uid,
		Weigh:              rule.Weigh,
	}, nil
}
