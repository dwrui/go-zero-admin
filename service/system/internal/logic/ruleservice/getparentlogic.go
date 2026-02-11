package ruleservicelogic

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gconv"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParentLogic {
	return &GetParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetParentLogic) GetParent(in *system.GetRuleParentRequest) (*system.GetRuleParentResponse, error) {
	ruleList, err := model.GetParentAll(l.ctx, l.svcCtx, ga.Slice{0, 1}, ga.Map{"id != ?": in.Id}, "id,pid,type,title,locale,icon,permission,path,component,weigh,status,create_time", "weigh")
	if err != nil {
		return nil, err
	}
	if len(ruleList) == 0 {
		return &system.GetRuleParentResponse{}, nil
	}
	newList := make([]map[string]interface{}, 0)
	for _, val := range ruleList {
		if val.Title == "" {
			val.Title = val.Locale
		}
		newList = append(newList, gconv.Map(val))
	}
	menuLists := ga.GetMenuChildrenArray(newList, 0, "pid")
	menuListJson, _ := json.Marshal(menuLists)
	return &system.GetRuleParentResponse{
		Data: ga.String(menuListJson),
	}, nil
}
