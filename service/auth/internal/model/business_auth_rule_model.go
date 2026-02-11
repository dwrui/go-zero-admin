package model

import "database/sql"

type BusinessAuthRuleModel struct {
	Id                 uint64         `db:"id"`
	Uid                int64          `db:"uid"`                // 添加用户
	Title              string         `db:"title"`              // 菜单名称
	Des                string         `db:"des"`                // 描述
	Locale             string         `db:"locale"`             // 中英文标题key
	Weigh              int64          `db:"weigh"`              // 排序
	Type               int64          `db:"type"`               // 类型 0=目录，1=菜单，2=按钮
	Pid                int64          `db:"pid"`                // 上一级
	Icon               string         `db:"icon"`               // 图标
	Routepath          string         `db:"routepath"`          // 路由地址
	Routename          string         `db:"routename"`          // 路由名称
	Component          string         `db:"component"`          // 组件路径
	Redirect           string         `db:"redirect"`           // 重定向地址
	Path               string         `db:"path"`               // 接口路径
	Permission         sql.NullString `db:"permission"`         // 权限标识前端验证
	ApiAuth            string         `db:"api_auth"`           // 权限标识后端验证
	Status             int64          `db:"status"`             // 状态 0=启用1=禁用
	Isext              int64          `db:"isext"`              // 是否外链 0=否1=是
	Keepalive          int64          `db:"keepalive"`          // 是否缓存 0=否1=是
	Requiresauth       int64          `db:"requiresauth"`       // 是否需要登录鉴权 0=否1=是
	Hideinmenu         int64          `db:"hideinmenu"`         // 是否在左侧菜单中隐藏该项 0=否1=是
	Hidechildreninmenu int64          `db:"hidechildreninmenu"` // 强制在左侧菜单中显示单项 0=否1=是
	Activemenu         int64          `db:"activemenu"`         // 高亮设置的菜单项 0=否1=是
	Noaffix            int64          `db:"noaffix"`            // 如果设置为true，标签将不会添加到tab-bar中 0=否1=是
	Onlypage           int64          `db:"onlypage"`           // 独立页面不需layout和登录，如登录页、数据大屏
	Createtime         sql.NullTime   `db:"createtime"`         // 创建时间
}
