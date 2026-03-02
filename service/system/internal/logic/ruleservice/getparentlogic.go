package ruleservicelogic

import (
	"context"
	"system/internal/model"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"

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
		newList = append(newList, ga.Map{
			"id":          val.Id,
			"title":       val.Title,
			"locale":      val.Locale,
			"icon":        val.Icon,
			"permission":  val.Permission,
			"path":        val.Path,
			"component":   val.Component,
			"weigh":       val.Weigh,
			"status":      val.Status,
			"create_time": val.CreateTime.Time.Format("2006-01-02 15:04:05"), // 处理sql.NullTime类型
			"type":        val.Type,
			"uid":         val.Uid,
			"pid":         val.Pid,
		})
	}
	menuLists := ga.GetMenuChildrenArray(newList, 0, "pid")
	list := ga.Map{
		"tree": menuLists,
		"list": newList,
	}
	menuListJson, _ := json.Marshal(list)
	return &system.GetRuleParentResponse{
		Data: ga.String(menuListJson),
	}, nil
}
