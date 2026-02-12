package model

import (
	"context"
	"database/sql"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"system/internal/svc"
	"system/system"
)

type AdminAuthDeptModel struct {
	Id         uint64       `db:"id"`
	BusinessId int64        `db:"business_id"` // 业务主账号id
	AccountId  int64        `db:"account_id"`  // 添加账号
	Name       string       `db:"name"`        // 部门名称
	Pid        int64        `db:"pid"`         // 上级部门
	Weigh      int64        `db:"weigh"`       // 排序
	Status     int64        `db:"status"`      // 状态
	Remark     string       `db:"remark"`      // 备注
	CreateTime sql.NullTime `db:"create_time"` // 创建时间
}

func GetDeptList(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetDeptListRequest) ([]*AdminAuthDeptModel, error) {
	whereMap := gmap.New()
	if req.BusinessId != 0 {
		whereMap.Set("business_id =?", req.BusinessId)
	}
	if req.Name != "" {
		whereMap.Set("name like ?", "%"+req.Name+"%")
	}
	if req.Status != 0 {
		whereMap.Set("status =?", req.Status)
	}
	if req.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(ga.String(req.CreateTime), ",")
		whereMap.Set("createtime between ? and ?", ga.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	var list []*AdminAuthDeptModel
	resp := svcCtx.DB.Model("admin_auth_dept").Where(whereMap).OrderBy("weigh").Select(ctx, &list)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	return list, nil
}
func GetDeptParent(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetDeptParentRequest) ([]*AdminAuthDeptModel, error) {
	whereMap := gmap.New()
	if req.BusinessId != 0 {
		whereMap.Set("business_id =?", req.BusinessId)
	}
	whereMap.Set("status =?", 0)
	var list []*AdminAuthDeptModel
	resp := svcCtx.DB.Model("admin_auth_dept").Fields("id,pid,name").Where(whereMap).OrderBy("weigh").Select(ctx, &list)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	return list, nil
}

func SaveDept(ctx context.Context, svcCtx *svc.ServiceContext, req *system.SaveDeptRequest) (uint64, error) {
	resp := svcCtx.DB.Model("admin_auth_dept").Data(req).Save(ctx)
	if resp.GetError() != nil {
		return 0, resp.GetError()
	}
	return ga.Uint64(resp.GetLastId()), nil
}
func UpStatusDept(ctx context.Context, svcCtx *svc.ServiceContext, req *system.UpStatusDeptRequest) (uint64, error) {
	resp := svcCtx.DB.Model("admin_auth_dept").Where("id =?", req.Id).Data(map[string]interface{}{"status": req.Status}).Update(ctx)
	if resp.GetError() != nil {
		return 0, resp.GetError()
	}
	return ga.Uint64(resp.GetLastId()), nil
}
func DelDept(ctx context.Context, svcCtx *svc.ServiceContext, req *system.DelDeptRequest) error {
	resp := svcCtx.DB.Model("admin_auth_dept").Where("id =?", req.Id).Delete(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}
