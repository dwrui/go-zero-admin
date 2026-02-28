package ruleservicelogic

import (
	"context"
	"fmt"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
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
		// 手动创建map，确保所有字段都能正确访问
		item := map[string]interface{}{
			"id":          val.Id,
			"pid":         val.Pid,
			"type":        val.Type,
			"title":       val.Title,
			"locale":      val.Locale,
			"icon":        val.Icon,
			"permission":  val.Permission.String, // 处理sql.NullString类型
			"path":        val.Path,
			"component":   val.Component,
			"weigh":       val.Weigh,
			"status":      val.Status,
			"create_time": val.Createtime.Time.Format("2006-01-02 15:04:05"), // 处理sql.NullTime类型
		}
		newList = append(newList, item)
	}

	menuLists := ga.GetTreeArray(newList, 0, "")
	// 递归函数，将map转换为RuleListData结构
	var convertToRuleListData func(menuList map[string]interface{}) *system.RuleListData
	convertToRuleListData = func(menuList map[string]interface{}) *system.RuleListData {
		// 处理children字段，从gvar.Var中提取实际数据
		var children []*system.RuleListData
		if childrenData := menuList["children"]; childrenData != nil {
			// 转换为[]map[string]interface{}
			if childrenMaps, ok := childrenData.([]map[string]interface{}); ok {
				for _, childMap := range childrenMaps {
					children = append(children, convertToRuleListData(childMap))
				}
			}
		}

		return &system.RuleListData{
			Id:         ga.Uint64(menuList["id"]),
			Pid:        ga.Uint64(menuList["pid"]),
			Type:       ga.Uint64(menuList["type"]),
			Title:      ga.String(menuList["title"]),
			Locale:     ga.String(menuList["locale"]),
			Icon:       ga.String(menuList["icon"]),
			Permission: ga.String(menuList["permission"]),
			Path:       ga.String(menuList["path"]),
			Component:  ga.String(menuList["component"]),
			Spacer:     ga.String(menuList["spacer"]),
			Weigh:      ga.Uint64(menuList["weigh"]),
			Status:     ga.Uint64(menuList["status"]),
			CreateTime: ga.String(menuList["create_time"]),
			Children:   children,
		}
	}
	menuListEnd := make([]*system.RuleListData, 0)
	for _, menuList := range menuLists {
		menuListEnd = append(menuListEnd, convertToRuleListData(menuList))
	}
	fmt.Println("menuList", menuListEnd)
	return &system.GetRuleListResponse{
		Data: menuListEnd,
	}, nil
}
