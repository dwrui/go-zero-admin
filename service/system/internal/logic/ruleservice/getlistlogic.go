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

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *system.GetRuleListRequest) (*system.GetRuleListResponse, error) {
	ruleList, err := model.GetRuleAll(l.ctx, l.svcCtx, "id,pid,type,title,locale,icon,permission,path,component,weigh,status,create_time", "weigh")
	if err != nil {
		return nil, err
	}
	if len(ruleList) == 0 {
		return &system.GetRuleListResponse{}, nil
	}
	newList := make([]map[string]interface{}, 0)
	for _, val := range ruleList {
		if val.Title == "" {
			val.Title = val.Locale
		}
		newList = append(newList, gconv.Map(val))
	}
	menuLists := ga.GetTreeArray(newList, 0, "")
	menuListJson, _ := json.Marshal(menuLists)
	return &system.GetRuleListResponse{
		Data: ga.String(menuListJson),
	}, nil
}
