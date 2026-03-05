package accountservicelogic

import (
	"context"
	"system/internal/model"
	"system/internal/svc"
	"system/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleLogic) GetRole(in *system.GetAccountRoleRequest) (*system.GetAccountRoleResponse, error) {
	resp, err := model.GetRole(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	// 递归函数，将map转换为RuleListData结构
	var convertToRuleListData func(list map[string]interface{}) *system.RoleData
	convertToRuleListData = func(list map[string]interface{}) *system.RoleData {
		// 处理children字段，从gvar.Var中提取实际数据
		var children []*system.RoleData
		if childrenData := list["children"]; childrenData != nil {
			// 转换为[]map[string]interface{}
			if childrenMaps, ok := childrenData.([]map[string]interface{}); ok {
				for _, childMap := range childrenMaps {
					children = append(children, convertToRuleListData(childMap))
				}
			}
		}

		return &system.RoleData{
			Id:         ga.Uint64(list["id"]),
			Pid:        ga.Uint64(list["pid"]),
			Name:       ga.String(list["name"]),
			Rules:      ga.String(list["rules"]),
			Menu:       ga.String(list["menu"]),
			Btns:       ga.String(list["btns"]),
			DataAccess: ga.Uint64(list["data_access"]),
			Remark:     ga.String(list["remark"]),
			AccountId:  ga.Uint64(list["account_id"]),
			BusinessId: ga.Uint64(list["business_id"]),
			Spacer:     ga.String(list["spacer"]),
			Weigh:      ga.Uint64(list["weigh"]),
			Status:     ga.Uint64(list["status"]),
			CreateTime: ga.String(list["create_time"]),
			Children:   children,
		}
	}
	menuListEnd := make([]*system.RoleData, 0)
	for _, list := range resp {
		menuListEnd = append(menuListEnd, convertToRuleListData(list))
	}
	return &system.GetAccountRoleResponse{
		List: menuListEnd,
	}, nil
}
