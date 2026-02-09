package model

import (
	"auth/internal/svc"
	"context"
	"errors"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"
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

/**
 * 设置用户权限
 * @param ctx 上下文
 * @param svg 服务上下文
 * @param userId 用户id
 * @param businessId 业务id
 */
func SetUserPermission(ctx context.Context, svg *svc.ServiceContext, userId uint64, businessId uint64) (map[string]interface{}, error) {
	var permissions = make(map[string]interface{})
	var BusinessAuthRole BusinessAuthRoleModel
	resp := svg.DB.Model("business_auth_role").Where(ga.Map{"business_id": businessId, "account_id": userId}).Find(ctx, &BusinessAuthRole)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	//权限为空返回错误
	if BusinessAuthRole.Rules == "" {
		return nil, errors.New("用户角色暂无权限")
	}
	permissions["roles"] = ga.String(BusinessAuthRole.Rules)
	if BusinessAuthRole.Rules != "*" {
		roles := svg.DB.Model("business_auth_rule").WhereIn("id", ga.Axplode(BusinessAuthRole.Rules)).Where("api_auth != ?", "").Where("type != ?", 0).Column(ctx, "api_auth")
		if roles.GetError() != nil || roles.IsEmpty() {
			return nil, errors.New("用户角色暂无权限")
		}
		permissions["permissions"] = ga.FormatColumnData(roles.GetData())
	} else {
		permissions["permissions"] = []string{"*"}
	}
	dataJson, err := json.Marshal(permissions)
	if err != nil {
		return nil, errors.New("用户权限错误")
	}
	//缓存24小时
	err = svg.Redis.SetexCtx(ctx, "user_permissions:"+ga.String(userId), ga.String(dataJson), 86400)
	if err != nil {
		return nil, errors.New("用户权限错误")
	}
	return permissions, nil
}
