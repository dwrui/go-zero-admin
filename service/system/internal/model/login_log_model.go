package model

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"net"
	"system/internal/svc"
	"system/system"
	"time"
)

type LoginLogModel struct {
	Id         uint64            `db:"id"`
	Uid        uint64            `db:"uid"`
	AccountId  uint64            `db:"account_id"`
	BusinessId uint64            `db:"business_id"`
	Type       string            `db:"type"`
	Status     int               `db:"status"` // 0:失败, 1:成功
	Des        string            `db:"des"`
	Ip         string            `db:"ip"`
	Address    string            `db:"address"`
	UserAgent  string            `db:"user_agent"`
	ErrorMsg   string            `db:"error_msg"`
	CreateTime time.Time         `db:"create_time"`
	User       map[string]string `db:"-"`
}

// 临时结构体：用于接收JOIN查询结果
type LogWithUserInfo struct {
	LoginLogModel
	UserAvatar   string `db:"user_avatar"`
	UserName     string `db:"user_name"`
	UserNickname string `db:"user_nickname"`
	UserUsername string `db:"user_username"`
}

func GetLoginLogList(ctx context.Context, svcCtx *svc.ServiceContext, req *system.GetLogListRequest) (ga.Map, error) {
	whereMap := gmap.New()
	whereMap.Set("log.business_id", req.BusinessId)

	if req.User != "" {
		userids := svcCtx.DB.Model("admin_account").Where("name like ?", "%"+ga.String(req.User)+"%").Column(ctx, "id")
		whereMap.Set("log.uid IN(?)", userids.GetData())
	}
	if req.Ip != "" {
		address := net.ParseIP(ga.String(req.Ip))
		if address == nil {
			whereMap.Set("log.address like ?", "%"+ga.String(address)+"%")
		} else {
			whereMap.Set("log.ip", req.Ip)
		}
	}
	if req.Status != 0 {
		whereMap.Set("log.status", req.Status)
	}
	if req.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(ga.String(req.CreateTime), ",")
		whereMap.Set("log.create_time between ? and ?", ga.Slice{datetime_arr[0], datetime_arr[1]})
	}
	var logsWithUser []LogWithUserInfo
	list := svcCtx.DB.Model("common_sys_login_log").Alias("log").Fields("log.*,ba.avatar as user_avatar,ba.name as user_name,ba.nickname as user_nickname,ba.username as user_username").LeftJoin("admin_account", "ba", "log.uid = ba.id").Where("log.type", "business").Where(whereMap).OrderByDesc("log.id").Paginate(ctx, ga.Int(req.Page), ga.Int(req.PageSize), &logsWithUser)
	if list.Error != nil {
		return nil, list.Error
	}
	logList := make([]*LoginLogModel, len(logsWithUser))
	if list.Total > 0 {
		for i, value := range logsWithUser {
			logItem := value.LoginLogModel
			logItem.User = ga.MapStrStr{
				"avatar":   value.UserAvatar,
				"name":     value.UserName,
				"nickname": value.UserNickname,
				"username": value.UserUsername,
			}
			logList[i] = &logItem
		}
	}
	return ga.Map{"items": logList, "total": list.Total, "page": req.Page, "page_size": req.PageSize}, nil
}
