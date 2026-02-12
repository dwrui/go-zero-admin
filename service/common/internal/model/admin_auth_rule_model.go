package model

import (
	"common/internal/svc"
	"context"
	"database/sql"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

type AdminAuthRule struct {
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
	Permission         sql.NullString `db:"permission"`         // 权限标识
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

// 获取权限菜单
func GetMenuArray(ctx context.Context, svg *svc.ServiceContext, pdata []AdminAuthRule, parent_id int64, roles []interface{}) []map[string]interface{} {
	var returnList []map[string]interface{}
	var one int64 = 1
	for _, v := range pdata {
		if ga.Int64(v.Pid) == parent_id {
			mid_item := map[string]interface{}{
				"path":      v.Routepath,
				"name":      v.Routename,
				"component": v.Component,
			}
			children := GetMenuArray(ctx, svg, pdata, ga.Int64(v.Id), roles)
			if children != nil {
				mid_item["children"] = children
			}
			//1.标题
			// var Menu_title interface{}
			// if v["locale"] != nil && v["locale"].String() != "" {
			// 	Menu_title = v["locale"]
			// } else {
			// 	Menu_title = v["title"]
			// }
			meta := map[string]interface{}{
				"locale": v.Locale,
				"title":  v.Title,
				"id":     v.Id,
			}
			//2.重定向
			if v.Redirect != "" {
				mid_item["redirect"] = v.Redirect
			}
			//3.隐藏子菜单
			if v.Hidechildreninmenu == one {
				meta["hideChildrenInMenu"] = true
			}
			//3.图标
			if v.Icon != "" {
				meta["icon"] = v.Icon
			}
			//4.缓存
			if v.Keepalive == one { //设置为true页面将不会被缓存 false=缓存
				meta["ignoreCache"] = false
			} else {
				meta["ignoreCache"] = true
			}
			//5.隐藏菜单
			if v.Hideinmenu == one {
				meta["hideInMenu"] = true
			}
			//6.在标签隐藏
			if v.Noaffix == one {
				meta["noAffix"] = true
			}
			//7.详情页在本业打开-用于配置详情页时左侧激活的菜单路径
			if v.Activemenu == one {
				meta["activeMenu"] = true
			}
			//8.是否需要登录鉴权
			if v.Requiresauth == one {
				meta["requiresAuth"] = true
			} else {
				meta["requiresAuth"] = false
			}
			//9.是否需要登录鉴权
			if v.Isext == one {
				meta["isExt"] = true
			}
			//10.是否需要登录鉴权
			if v.Onlypage == one {
				meta["onlypage"] = true
			} else {
				meta["onlypage"] = false
			}
			//11.按钮权限
			if len(roles) == 0 { //超级权限
				permission := svg.DB.Model("admin_auth_rule").Where("status", 0).Where("type", 2).Where("pid", v.Id).WhereNotNull("permission").Column(ctx, "permission")
				if permission.IsNotEmpty() && len(permission.GetData().([]string)) > 0 {
					meta["btnroles"] = permission.GetData().([]string)
				} else {
					meta["btnroles"] = [1]string{"*"}
				}
			} else { //选择路由
				permission := svg.DB.Model("admin_auth_rule").Where("status", 0).Where("type", 2).Where("pid", v.Id).WhereIn("id", roles).WhereNotNull("permission").Column(ctx, "permission")
				if permission.IsNotEmpty() && len(permission.GetData().([]string)) > 0 {
					meta["btnroles"] = permission.GetData().([]string)
				} else {
					hasepermission := svg.DB.Model("admin_auth_rule").Where("status", 0).Where("type", 2).Where("pid", v.Id).WhereNotNull("permission").Column(ctx, "permission")
					if hasepermission.IsEmpty() {
						meta["btnroles"] = make([]interface{}, 0)
					} else {
						meta["btnroles"] = [1]string{"*"}
					}
				}
			}
			//赋值
			mid_item["meta"] = meta
			returnList = append(returnList, mid_item)
		}
	}
	return returnList
}
