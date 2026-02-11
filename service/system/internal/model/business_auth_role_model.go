package model

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"system/internal/svc"
	"system/system"
)

type BusinessAuthRoleModel struct {
	Id         uint64 `db:"id"`
	BusinessId int64  `db:"business_id"` // 业务主账号id
	AccountId  int64  `db:"account_id"`  // 添加用户id
	Pid        int64  `db:"pid"`         // 父级
	Name       string `db:"name"`        // 名称
	Rules      string `db:"rules"`       // 规则ID 所拥有的权限包扣父级
	Menu       string `db:"menu"`        // 选择的id，用于编辑赋值
	Btns       string `db:"btns"`        // 按钮id，用于编辑赋值
	Status     int64  `db:"status"`      // 状态1=禁用
	DataAccess int64  `db:"data_access"` // 数据权限0=自己1=自己及子权限，2=全部
	Remark     string `db:"remark"`      // 描述
	Weigh      int64  `db:"weigh"`       // 排序
}

func GetRoleList(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetRoleListRequest) ([]*BusinessAuthRoleModel, error) {
	user_role_ids := svcCtx.DB.Model("business_auth_role_access").Where("uid = ?", req.UserId).Column(ctx, "role_id")
	var allRoleModel []*BusinessAuthRoleModel
	allRole := svcCtx.DB.Model("business_auth_role").All(ctx, &allRoleModel)
	if allRole.GetError() != nil {
		return nil, allRole.GetError()
	}
	allRoleMap := make([]map[string]interface{}, 0)
	for _, v := range allRoleModel {
		allRoleMap = append(allRoleMap, ga.Map{
			"id":  v.Id,
			"pid": v.Pid,
		})
	}
	role_chil_ids := ga.FindAllChildrenIDs(allRoleMap, user_role_ids.GetData().([]uint64)) //批量获取子节点id
	all_role_id := append(user_role_ids.GetData().([]uint64), role_chil_ids...)
	whereMap := gmap.New()
	whereMap.Set("id IN(?)", all_role_id) //in 查询
	return nil, nil
}
