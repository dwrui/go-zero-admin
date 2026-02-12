package model

import (
	"context"
	"fmt"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"net"
	"system/internal/svc"
	"system/system"
	"time"
)

type LoginLogModel struct {
	Id         uint64    `db:"id"`
	Uid        uint64    `db:"uid"`
	AccountId  uint64    `db:"account_id"`
	BusinessId uint64    `db:"business_id"`
	Type       string    `db:"type"`
	Status     int       `db:"status"` // 0:失败, 1:成功
	Des        string    `db:"des"`
	Ip         string    `db:"ip"`
	Address    string    `db:"address"`
	UserAgent  string    `db:"user_agent"`
	ErrorMsg   string    `db:"error_msg"`
	CreateTime time.Time `db:"create_time"`
}

func GetLoginLogList(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetLogListRequest) (ga.Map, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	whereMap := gmap.New()

	whereMap.Set("business_id", req.BusinessId)

	if req.User != "" {
		userids := svcCtx.DB.Model("business_account").Where("name like ?", "%"+ga.String(req.User)+"%").Column(ctx, "id")
		whereMap.Set("uid IN(?)", userids.GetData())
	}
	if req.Ip != "" {
		address := net.ParseIP(ga.String(req.Ip))
		if address == nil {
			whereMap.Set("address like ?", "%"+ga.String(address)+"%")
		} else {
			whereMap.Set("ip", req.Ip)
		}
	}
	if req.Status != 0 {
		whereMap.Set("status", req.Status)
	}
	if req.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(ga.String(req.CreateTime), ",")
		whereMap.Set("createtime between ? and ?", ga.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	var logList []*LoginLogModel
	list := svcCtx.DB.Model("common_sys_login_log").SQLFetch(true).Where("type", "business").Where(whereMap).OrderByDesc("id").Paginate(ctx, ga.Int(req.Page), ga.Int(req.PageSize), &logList)
	fmt.Println(logList)
	if list.Error != nil {
		return nil, list.Error
	}
	for _, vv := range logList {
		fmt.Println(vv)
	}
	return ga.Map{"items": logList, "total": list.Total, "page": req.Page, "page_size": req.PageSize}, nil
}
