package model

import (
	"context"
	"database/sql"
	"errors"
	"system/internal/svc"
	"system/system"
	"time"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
)

type AdminAccountModel struct {
	Id            uint64       `db:"id"`
	AccountId     int64        `db:"account_id"`      // 账号id/记录那个账号添加
	BusinessId    int64        `db:"business_id"`     // 业务主账号id
	MainAccount   int64        `db:"main_account"`    // 是否为主账号:0=否,1=是
	DeptId        int64        `db:"dept_id"`         // 部门id
	Username      string       `db:"username"`        // 用户账号
	Password      string       `db:"password"`        // 密码
	Salt          string       `db:"salt"`            // 密码盐
	Name          string       `db:"name"`            // 姓名
	Nickname      string       `db:"nickname"`        // 昵称
	Avatar        string       `db:"avatar"`          // 头像
	Tel           string       `db:"tel"`             // 备用电话用户自己填写
	Mobile        string       `db:"mobile"`          // 手机号码
	Email         string       `db:"email"`           // 邮箱
	LastLoginIp   string       `db:"last_login_ip"`   // 最后登录IP
	LastLoginTime int64        `db:"last_login_time"` // 最后登录时间
	Status        int64        `db:"status"`          // 状态0=正常，1=禁用
	Validtime     int64        `db:"validtime"`       // 账号有效时间0=无限
	Address       string       `db:"address"`         // 地址
	Remark        string       `db:"remark"`          // 描述
	Company       string       `db:"company"`         // 公司名称
	Province      string       `db:"province"`        // 省份
	City          string       `db:"city"`            // 城市
	Area          string       `db:"area"`            // 地区
	FileSize      uint64       `db:"fileSize"`        // 附件存储空间
	LoginAttempts int64        `db:"login_attempts"`  // 登录尝试次数
	LockTime      time.Time    `db:"lock_time"`       // 账号锁定时间
	CreateTime    sql.NullTime `db:"create_time"`     // 创建时间
	UpdateTime    sql.NullTime `db:"update_time"`     // 修改时间
	DeleteTime    sql.NullTime `db:"delete_time"`     // 删除时间
	PwdResetTime  sql.NullTime `db:"pwd_reset_time"`  // 修改密码时间
}
type AdminAccountModelResponse struct {
	*AdminAccountModel
	RoleId   []uint64 `db:"-" json:"role_id"`   // 角色id
	RoleName []string `db:"-" json:"role_name"` // 角色名称
	DeptName string   `db:"-" json:"dept_name"` // 部门名称
}

func GetAccountList(ctx context.Context, svcCtx *svc.ServiceContext, whereMap *gmap.Map, page, size uint64) (ga.Map, error) {
	var list []*AdminAccountModelResponse

	resp := svcCtx.DB.Model("admin_account").Alias("c").Fields("c.id,c.status,c.name,c.username,c.avatar,c.tel,c.mobile,c.email,c.dept_id,c.remark,c.city,c.address,c.company,c.create_time,d.name as dept_name").LeftJoin("admin_auth_dept", "d", "c.dept_id = d.id").Where(whereMap).Order("c.id", "desc").Paginate(ctx, ga.Int(page), ga.Int(size), &list)
	if resp.Error != nil {
		return nil, resp.Error
	}
	for _, v := range list {
		roleid := []uint64{}
		roleName := []string{}
		roleResp := svcCtx.DB.Model("admin_auth_role_access").Where("uid", v.Id).Column(ctx, "role_id", &roleid)
		if roleResp.IsNotEmpty() {
			svcCtx.DB.Model("admin_auth_role").WhereIn("id", roleid).Column(ctx, "name", &roleName)
		}
		v.RoleId = roleid
		v.RoleName = roleName
	}
	return ga.Map{
		"list":  list,
		"page":  resp.Page,
		"size":  resp.Size,
		"total": resp.Total,
	}, nil
}

func SaveAccount(ctx context.Context, svcCtx *svc.ServiceContext, data ga.Map) (uint64, error) {
	resp := svcCtx.DB.Model("admin_account").Save(ctx, data)
	if resp.GetError() != nil {
		return 0, resp.GetError()
	}
	return ga.Uint64(resp.GetLastId()), nil
}

func AppRoleAccess(ctx context.Context, svg *svc.ServiceContext, roleids []uint64, uid interface{}) {
	//批量提交
	svg.DB.Model("admin_auth_role_access").Where("uid", uid).Delete(ctx)
	save_arr := ga.List{}
	for _, val := range roleids {
		marr := map[string]interface{}{"uid": uid, "role_id": val}
		save_arr = append(save_arr, marr)
	}
	svg.DB.Model("admin_auth_role_access").Data(save_arr).InsertAll(ctx)
}
func UpStatusAccount(ctx context.Context, svg *svc.ServiceContext, id uint64, status uint64) error {
	resp := svg.DB.Model("admin_account").Where("id", id).Update(ctx, ga.Map{"status": status})
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

func DelAccount(ctx context.Context, svg *svc.ServiceContext, id uint64) error {
	resp := svg.DB.Model("admin_account").Where("id = ?", id).Delete(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

// 获取所有子级ID（包含自身）
func GetRole(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetAccountRoleRequest) (ga.List, error) {
	var user_role_ids []uint64
	svcCtx.DB.Model("admin_auth_role_access").Where("uid = ?", req.UserId).Column(ctx, "role_id", &user_role_ids)
	var allRoleModel []*AdminAuthRoleModel
	allRole := svcCtx.DB.Model("admin_auth_role").All(ctx, &allRoleModel)
	if allRole.GetError() != nil {
		return nil, allRole.GetError()
	}
	allRoleMap := make([]map[string]uint64, 0)
	for _, v := range allRoleModel {
		allRoleMap = append(allRoleMap, map[string]uint64{
			"id":  v.Id,
			"pid": ga.Uint64(v.Pid),
		})
	}
	role_chil_ids := ga.FindAllChildrenIDs(allRoleMap, user_role_ids) //批量获取子节点id
	all_role_id := append(user_role_ids, role_chil_ids...)
	whereMap := gmap.New()
	if len(all_role_id) > 0 {
		whereMap.Set("id IN(?)", all_role_id) //in 查询
	}
	account_id, _ := GetDataAuthor(ctx, svcCtx, req.UserId, "")
	account_id = append(account_id, 0)
	var my_role_account_id []uint64
	svcCtx.DB.Model("admin_auth_role").WhereIn("id", user_role_ids).Column(ctx, "account_id", &my_role_account_id)
	//合并account_id和myRoleIds
	account_ids := append(account_id, my_role_account_id...)
	whereMap.Set("account_id IN(?)", account_ids)
	var roleList []*AdminAuthRoleModel
	roleListData := svcCtx.DB.Model("admin_auth_role").Where(whereMap).OrderBy("weigh").Select(ctx, &roleList)
	if roleListData.GetError() != nil {
		return nil, roleListData.GetError()
	}
	//获取最大一级的pid
	max_role_id := svcCtx.DB.Model("admin_auth_role").Where(whereMap).OrderBy("id").Value(ctx, "pid")
	roleListMap := make([]map[string]interface{}, 0)
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

	roleListTree := ga.GetTreeArray(roleListMap, ga.Int64(max_role_id.GetData()), "")
	if roleListTree == nil {
		roleListTree = make([]map[string]interface{}, 0)
	}
	return roleListTree, nil
}

func Isaccountexist(ctx context.Context, svcCtx *svc.ServiceContext, id uint64, username string) error {
	if id == 0 {
		resp := svcCtx.DB.Model("admin_account").Where("username = ?", username).Value(ctx, "id")
		if resp.GetError() != nil {
			return resp.GetError()
		}
		if resp.IsNotEmpty() {
			return errors.New("用户名已存在")
		}
	} else {
		resp := svcCtx.DB.Model("admin_account").Where("username = ?", username).Where("id = ?", id).Value(ctx, "id")
		if resp.GetError() != nil {
			return resp.GetError()
		}
		if resp.IsNotEmpty() {
			return errors.New("用户名已存在")
		}
	}
	return nil
}
