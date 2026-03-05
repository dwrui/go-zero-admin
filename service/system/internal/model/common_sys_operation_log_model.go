package model

import (
	"context"
	"database/sql"
	"system/internal/svc"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gtime"
)

type CommonSysOperationLogMpdel struct {
	Id          uint64       `db:"id"`
	BusinessId  uint64       `db:"business_id"`  // 业务主账号id
	AccountId   uint64       `db:"account_id"`   // 添加用户id
	UserId      uint64       `db:"user_id"`      // 用户id
	Type        string       `db:"type"`         // 操作类型
	Method      string       `db:"method"`       // 操作方法
	Path        string       `db:"path"`         // 操作路径
	IP          string       `db:"ip"`           // 操作IP
	Address     string       `db:"address"`      // 操作地址
	Description string       `db:"description"`  // 操作描述
	ReqHeaders  string       `db:"req_headers"`  // 操作请求头
	ReqBody     string       `db:"req_body"`     // 操作请求体
	RespHeaders string       `db:"resp_headers"` // 操作响应头
	RespBody    string       `db:"resp_body"`    // 操作响应体
	Duration    float64      `db:"duration"`     // 操作响应时间
	Status      uint32       `db:"status"`       // 操作状态
	CreateTime  sql.NullTime `db:"create_time"`  // 创建时间
}
type LogOperationWithUserInfo struct {
	*CommonSysOperationLogMpdel
	Name         sql.NullString `db:"name"`
	UserName     sql.NullString `db:"user_name"`
	UserNickname sql.NullString `db:"user_nickname"`
	Avatar       sql.NullString `db:"avatar"`
}

// GetOperationLogList 获取操作日志列表
func GetOperationLogList(ctx context.Context, svcCtx *svc.ServiceContext, whereMap *gmap.Map, page, pageSize uint64) (ga.Map, error) {
	var list []*LogOperationWithUserInfo
	resp := svcCtx.DB.Model("common_sys_operation_log").Alias("log").Fields("log.id,log.user_id,log.type,log.method,log.ip,log.path,log.address,log.description,log.duration,log.status,log.create_time,u.name,u.avatar,u.nickname as user_nickname").LeftJoin("admin_account", "u", "log.account_id = u.id").Where(whereMap).OrderByDesc("log.id").Paginate(ctx, int(page), int(pageSize), &list)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return ga.Map{
		"items": list,
		"page":  page,
		"size":  pageSize,
		"total": resp.Total,
	}, nil
}

// GetOperatuinDetail 获取操作日志详情
func GetOperatuinDetail(ctx context.Context, svcCtx *svc.ServiceContext, id uint64) (LogOperationWithUserInfo, error) {
	var log LogOperationWithUserInfo
	resp := svcCtx.DB.Model("common_sys_operation_log").Alias("log").LeftJoin("admin_account", "u", "log.account_id = u.id").Fields("log.*,u.name as user_name").Where("log.id = ?", id).Find(ctx, &log)
	if resp.GetError() != nil {
		return LogOperationWithUserInfo{}, resp.GetError()
	}
	return log, nil
}

// DeleteOperationLog 删除1个月前的操作日志
func DeleteOperationLog(ctx context.Context, svcCtx *svc.ServiceContext) error {
	res := svcCtx.DB.Model("common_sys_operation_log").Where("create_time < ?", gtime.Now().AddDate(0, -1, 0).Format("Y-m-d H:i:s")).Delete(ctx)
	if res.GetError() != nil {
		return res.GetError()
	}
	return nil
}
