package model

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/plugin"
	"time"
	"user/internal/svc"
)

type LoginLogModel struct {
	Id          int64     `db:"id"`
	Uid         int64     `db:"uid"`
	AccountId   int64     `db:"account_id"`
	BusinessId  int64     `db:"business_id"`
	Type        string    `db:"type"`
	Status      int       `db:"status"` // 0:失败, 1:成功
	Des         string    `db:"des"`
	Ip          string    `db:"ip"`
	Address     string    `db:"address"`
	UserAgent   string    `db:"user_agent"`
	ErrorMsg    string    `db:"error_msg"`
	CreatedTime time.Time `db:"created_time"`
}

/**
 * 添加登录日志
 * @param ctx
 * @param svg
 * @param savedata
 */
func AddloginLog(ctx context.Context, svg *svc.ServiceContext, savedata ga.Map) {
	address, err := plugin.NewIpRegion(ga.String(savedata["ip"]))
	savedata["address"] = ""
	if err == nil {
		savedata["address"] = address
	}
	svg.DB.Model("common_sys_login_log").Insert(ctx, savedata)
}
