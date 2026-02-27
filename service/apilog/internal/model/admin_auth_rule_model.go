package model

import (
	"apilog/internal/svc"
	"context"
	"database/sql"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

type AdminAuthRuleModel struct {
	Id                 uint64         `db:"id"`
	Uid                uint64         `db:"uid"`                // 添加用户
	Title              string         `db:"title"`              // 菜单名称
	Des                string         `db:"des"`                // 描述
	Locale             string         `db:"locale"`             // 中英文标题key
	Weigh              uint64         `db:"weigh"`              // 排序
	Type               uint64         `db:"type"`               // 类型 0=目录，1=菜单，2=按钮
	Pid                uint64         `db:"pid"`                // 上一级
	Icon               string         `db:"icon"`               // 图标
	Routepath          string         `db:"routepath"`          // 路由地址
	Routename          string         `db:"routename"`          // 路由名称
	Component          string         `db:"component"`          // 组件路径
	Redirect           string         `db:"redirect"`           // 重定向地址
	Path               string         `db:"path"`               // 接口路径
	Permission         sql.NullString `db:"permission"`         // 权限标识前端验证
	ApiAuth            string         `db:"api_auth"`           // 权限标识后端验证
	Status             uint64         `db:"status"`             // 状态 0=启用1=禁用
	Isext              uint64         `db:"isext"`              // 是否外链 0=否1=是
	Keepalive          uint64         `db:"keepalive"`          // 是否缓存 0=否1=是
	Requiresauth       uint64         `db:"requiresauth"`       // 是否需要登录鉴权 0=否1=是
	Hideinmenu         uint64         `db:"hideinmenu"`         // 是否在左侧菜单中隐藏该项 0=否1=是
	Hidechildreninmenu uint64         `db:"hidechildreninmenu"` // 强制在左侧菜单中显示单项 0=否1=是
	Activemenu         uint64         `db:"activemenu"`         // 高亮设置的菜单项 0=否1=是
	Noaffix            uint64         `db:"noaffix"`            // 如果设置为true，标签将不会添加到tab-bar中 0=否1=是
	Onlypage           uint64         `db:"onlypage"`           // 独立页面不需layout和登录，如登录页、数据大屏
	Createtime         sql.NullTime   `db:"createtime"`         // 创建时间
}

func GetDes(ctx context.Context, svc *svc.ServiceContext, path string) string {
	// 从数据库中查询描述
	var rule AdminAuthRuleModel
	resp := svc.DB.Model("admin_auth_rule").Where("path = ?", path).Find(ctx, &rule)
	if resp.GetError() != nil || resp.IsEmpty() {
		return ""
	}
	p_title := svc.DB.Model("admin_auth_rule").Where("id = ?", rule.Pid).Value(ctx, "title")
	substr := rule.Des
	if rule.Des == "" {
		substr = rule.Title
	}
	return ga.String(p_title) + "【" + substr + "】"
}
