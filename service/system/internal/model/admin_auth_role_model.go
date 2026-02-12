package model

import (
	"context"
	"fmt"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gvar"
	"strings"
	"system/internal/svc"
	"system/system"
)

type AdminAuthRoleModel struct {
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

func GetRoleList(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetRoleListRequest) (ga.List, error) {
	user_role_ids := svcCtx.DB.Model("admin_auth_role_access").Where("uid = ?", req.UserId).Column(ctx, "role_id")
	var allRoleModel []*AdminAuthRoleModel
	allRole := svcCtx.DB.Model("admin_auth_role").All(ctx, &allRoleModel)
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
	account_id, _ := getDataAuthor(ctx, svcCtx, req.UserId, req.RequestUrl)
	account_id = append(account_id, 0)
	my_role_account_id := svcCtx.DB.Model("business_auth_role").WhereIn("id", user_role_ids.GetData().([]interface{})).Column(ctx, "account_id")
	account_id = append(account_id, my_role_account_id.GetData().([]uint64))
	whereMap.Set("account_id IN(?)", account_id)
	if req.Name != "" {
		whereMap.Set("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != 0 {
		whereMap.Set("status = ?", req.Status)
	}
	if req.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(req.CreateTime, ",")
		whereMap.Set("createtime between ? and ?", ga.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	var roleList []*AdminAuthRoleModel
	roleListData := svcCtx.DB.Model("business_auth_role").Where(whereMap).OrderBy("weigh").Select(ctx, &roleList)
	if roleListData.GetError() != nil {
		return nil, roleListData.GetError()
	}
	//获取最大一级的pid
	max_role_id := svcCtx.DB.Model("business_auth_role").Where(whereMap).OrderBy("id").Value(ctx, "pid")
	roleListMap := make([]map[string]interface{}, len(roleList))
	for _, val := range roleList {
		roleListMap = append(roleListMap, ga.Map{
			"id":          val.Id,
			"pid":         val.Pid,
			"name":        val.Name,
			"rules":       val.Rules,
			"menu":        val.Menu,
			"btns":        val.Btns,
			"status":      val.Status,
			"data_access": val.DataAccess,
			"remark":      val.Remark,
			"weigh":       val.Weigh,
			"business_id": val.BusinessId,
			"account_id":  val.AccountId,
		})
	}
	roleListTree := ga.GetTreeArray(roleListMap, ga.Int64(max_role_id), "")
	if roleListTree == nil {
		roleListTree = make([]map[string]interface{}, 0)
	}
	return roleListTree, nil
}

func getDataAuthor(ctx context.Context, svcCtx *svc.ServiceContext, userId uint64, requestUrl string) (ga.Slice, bool) {
	user_id := userId
	table_str := ""
	var acount_id = ga.Slice{user_id}
	if strings.HasPrefix(requestUrl, "/super-admin/") {
		table_str = "superadmin"
	} else if strings.HasPrefix(requestUrl, "/admin/") {
		table_str = "admin"
	}
	if table_str != "" {
		role_ids := svcCtx.DB.Model(table_str+"_auth_role_access").Where("uid", user_id).Column(ctx, "role_id")
		data_access := svcCtx.DB.Model(table_str+"_auth_role").WhereIn("id", role_ids.GetData().([]interface{})).Column(ctx, "data_access")
		data_access_info := data_access.GetData().([]*gvar.Var)
		if ga.IntInVarArray(1, data_access_info) { //数据权限0=自己1=自己及子权限，2=全部
			var allRuleModel []*AdminAuthRoleModel
			allRule := svcCtx.DB.Model(table_str+"_auth_rule").All(ctx, &allRuleModel)
			if allRule.GetError() != nil {
				return nil, false
			}
			allRuleMap := make([]map[string]interface{}, 0)
			for _, v := range allRuleModel {
				allRuleMap = append(allRuleMap, ga.Map{
					"id":  v.Id,
					"pid": v.Pid,
				})
			}
			chri_role_ids := ga.FindAllChildrenIDs(allRuleMap, role_ids.GetData().([]uint64)) //批量获取子节点id
			roleIdsInterface := make([]interface{}, len(chri_role_ids))
			for i, id := range chri_role_ids {
				roleIdsInterface[i] = id
			}
			uid_ids := svcCtx.DB.Model(table_str+"_auth_role_access").WhereIn("role_id", roleIdsInterface).Column(ctx, "uid")
			for _, val := range uid_ids.GetData().([]uint64) {
				acount_id = append(acount_id, val)
			}
			return acount_id, true //自己及子权限
		} else if ga.IntInVarArray(0, data_access_info) {
			return acount_id, true //自己
		}
	}
	return acount_id, false //全部
}

func GetRoleParent(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetRoleParentRequest) (ga.List, error) {
	user_role_ids := svcCtx.DB.Model("admin_auth_role_access").Where("uid = ?", req.Id).Column(ctx, "role_id")
	var allRoleModel []*AdminAuthRoleModel
	allRole := svcCtx.DB.Model("admin_auth_role").All(ctx, &allRoleModel)
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
	account_id, _ := getDataAuthor(ctx, svcCtx, req.UserId, req.RequestUrl)
	account_id = append(account_id, 0)
	my_role_account_id := svcCtx.DB.Model("business_auth_role").WhereIn("id", user_role_ids.GetData().([]interface{})).Column(ctx, "account_id")
	account_id = append(account_id, my_role_account_id.GetData().([]uint64))
	whereMap.Set("account_id IN(?)", account_id)
	whereMap.Set("id = ?", req.Id) //id 查询
	var roleList []*AdminAuthRoleModel
	roleListData := svcCtx.DB.Model("business_auth_role").Where(whereMap).OrderBy("weigh").Select(ctx, &roleList)
	if roleListData.GetError() != nil {
		return nil, roleListData.GetError()
	}
	//获取最大一级的pid
	max_role_id := svcCtx.DB.Model("business_auth_role").Where(whereMap).OrderBy("id").Value(ctx, "pid")
	roleListMap := make([]map[string]interface{}, len(roleList))
	for _, val := range roleList {
		roleListMap = append(roleListMap, ga.Map{
			"id":          val.Id,
			"pid":         val.Pid,
			"name":        val.Name,
			"rules":       val.Rules,
			"menu":        val.Menu,
			"btns":        val.Btns,
			"status":      val.Status,
			"data_access": val.DataAccess,
			"remark":      val.Remark,
			"weigh":       val.Weigh,
			"business_id": val.BusinessId,
			"account_id":  val.AccountId,
		})
	}
	roleListTree := ga.GetTreeArray(roleListMap, ga.Int64(max_role_id), "")
	if roleListTree == nil {
		roleListTree = make([]map[string]interface{}, 0)
	}
	return roleListTree, nil

}

func GetRoleMenuList(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetMenuListRequest) (ga.List, error) {
	var rule_ids []interface{}
	MDB := svcCtx.DB.Model("business_auth_rule").Where("status", 0).WhereIn("type", []interface{}{0, 1})
	if req.Pid == 0 {
		role_id := svcCtx.DB.Model("business_auth_role_access").Where("uid", req.UserId).Column(ctx, "role_id")
		menu_id := svcCtx.DB.Model("business_auth_role").WhereIn("id", role_id.GetData()).Column(ctx, "rules")
		//获取超级角色
		super_role := svcCtx.DB.Model("business_auth_role").WhereIn("id", role_id).Where("rules", "*").Value(ctx, "id")
		if super_role == nil { //不是超级权限-过滤菜单权限
			getmenus := ga.ArrayMerge(menu_id.GetData().([]*gvar.Var))
			MDB = MDB.WhereIn("id", getmenus)
			rule_ids = getmenus
		}
	} else {
		//获取用户权限
		menu_id_str := svcCtx.DB.Model("business_auth_role").Where("id", req.Pid).Value(ctx, "rules")
		if !strings.Contains(ga.String(menu_id_str.GetData()), "*") { //不是超级权限-过滤菜单权限
			getmenus := ga.Axplode(ga.String(menu_id_str.GetData()))
			MDB = MDB.WhereIn("id", getmenus)
			rule_ids = getmenus
		}
	}
	var menuListModel []*AdminAuthRuleModel
	menuList := MDB.Fields("id,pid,title,locale").OrderBy("weigh").Select(ctx, &menuListModel)
	if menuList.GetError() != nil {
		return nil, menuList.GetError()
	}
	menuListMap := make([]map[string]interface{}, len(menuListModel))
	for _, v := range menuListModel {
		menuListMap = append(menuListMap, ga.Map{
			"id":     v.Id,
			"pid":    v.Pid,
			"title":  v.Title,
			"locale": v.Locale,
		})
	}
	for _, val := range menuListMap {
		if val["title"] == "" {
			val["title"] = val["locale"]
		}
		delete(val, "locale")
		//获取按钮
		whereMap := gmap.New()
		if rule_ids != nil {
			whereMap.Set("id IN(?)", rule_ids)
		}
		var childrenMenuList []*AdminAuthRuleModel
		btn_rules := svcCtx.DB.Model("business_auth_rule").Where("status", 0).Where("type", 2).Where("pid", val["id"]).Where(whereMap).Fields("id,pid,title,des,locale").OrderBy("weigh").Select(ctx, &childrenMenuList)
		if btn_rules.IsNotEmpty() {
			item := ga.Map{
				"title":     "按钮权限",
				"id":        childrenMenuList[0].Id,
				"pid":       val["id"],
				"checkable": false,
				"btn_rules": btn_rules,
			}
			var valitem []ga.Map
			valitem = append(valitem, item)
			val["children"] = gvar.New(valitem)
			var btnids []interface{}
			for _, btnid := range childrenMenuList {
				btnids = append(btnids, btnid.Id)
			}
			val["btnids"] = gvar.New(btnids)
		} else if val["pid"] == 0 {
			//一级菜单获取子级菜单按钮
			sub_rule_ids := svcCtx.DB.Model("business_auth_rule").Where("pid", val["id"]).Where("status", 0).Where("type !=", 2).Column(ctx, "id")
			btn_rule_ids := svcCtx.DB.Model("business_auth_rule").Where("status", 0).Where("type", 2).WhereIn("pid", sub_rule_ids.GetData()).Column(ctx, "id")
			val["btnids"] = gvar.New(btn_rule_ids)
		}
		val["checkable"] = gvar.New(true)
	}
	menuTreeList := ga.GetMenuChildrenArray(menuListMap, 0, "pid")
	return menuTreeList, nil
}
func SaveRole(ctx context.Context, svcCtx *svc.ServiceContext, req *system.SaveRoleRequest) (uint64, error) {
	if req.Menu != "" && req.Menu != "*" {
		rules := GetRulesID("business_auth_rule", "pid", req.Menu, ctx, svcCtx) //获取子菜单包含的父级ID
		rudata := rules.([]interface{})
		var rulesStr []string
		for _, v := range rudata {
			str := fmt.Sprintf("%v", v) //interface{}强转string
			rulesStr = append(rulesStr, str)
		}
		for _, bv := range req.Btns {
			str := fmt.Sprintf("%v", bv) //interface{}强转string
			rulesStr = append(rulesStr, str)
		}
		req.Rules = strings.Join(rulesStr, ",")
	}

	resp := svcCtx.DB.Model("business_auth_role").Save(ctx, req)
	if resp.GetError() != nil {
		return 0, resp.GetError()
	}
	return ga.Uint64(resp.GetLastId()), nil

}
func UpStatusRole(ctx context.Context, svcCtx *svc.ServiceContext, req *system.UpStatusRoleRequest) error {
	resp := svcCtx.DB.Model("business_auth_role").Where("id", req.Id).Update(ctx, ga.Map{"status": req.Status})
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}
func DeleteRole(ctx context.Context, svcCtx *svc.ServiceContext, req *system.DelRoleRequest) error {

	resp := svcCtx.DB.Model("business_auth_role").Where("id", req.Id).Delete(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

// 获取子菜单包含的父级ID-返回全部ID
func GetRulesID(tablename string, field string, menus interface{}, ctx context.Context, svcCtx *svc.ServiceContext) interface{} {
	menus_rang := menus.([]interface{})
	var fnemuid []interface{}
	for _, v := range menus_rang {
		fid := getParentID(tablename, field, v, ctx, svcCtx)
		if fid != nil {
			fnemuid = ga.MergeArr_interface(fnemuid, fid)
		}
	}
	r_nemu := ga.MergeArr_interface(menus_rang, fnemuid)
	uni_fnemuid := ga.UniqueArr(r_nemu) //去重
	return uni_fnemuid
}

// 获取所有父级ID
func getParentID(tablename string, field string, id interface{}, ctx context.Context, svcCtx *svc.ServiceContext) []interface{} {
	var pids []interface{}
	pid := svcCtx.DB.Model(tablename).Where("id", id).Value(ctx, field)
	if pid != nil {
		a_pid := ga.Int64(pid.GetData())
		var zr_pid int64 = 0
		if a_pid != zr_pid {
			pids = append(pids, a_pid)
			getParentID(tablename, field, ga.Int64(a_pid), ctx, svcCtx)
		}
	}
	return pids
}
