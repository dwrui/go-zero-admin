package model

import (
	"apilog/apilog"
	"apilog/internal/svc"
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationLog struct {
	UserID      int64        `db:"user_id"`      //用户ID
	AccountID   int64        `db:"account_id"`   //账号ID
	BusinessID  int64        `db:"business_id"`  //业务ID
	Type        string       `db:"type"`         //日志类型 admin后台日志 adminpro总后台日志
	Method      string       `db:"method"`       //请求方法
	Path        string       `db:"path"`         //请求路径
	IP          string       `db:"ip"`           //请求IP
	Address     string       `db:"address"`      //根据ip获取的地址
	ReqHeaders  string       `db:"req_headers"`  //请求头
	ReqBody     string       `db:"req_body"`     //请求体
	RespHeaders string       `db:"resp_headers"` //响应头
	RespBody    string       `db:"resp_body"`    //响应体
	Status      int          `db:"status"`       //1成功0失败
	Duration    int64        `db:"duration"`     // 耗时
	CreatedTime sql.NullTime `db:"created_time"` // 创建时间
}

func AddOperationLog(ctx context.Context, svg *svc.ServiceContext, in *apilog.OperationLogRequest) error {
	resp := svg.DB.Model("common_sys_operation_log").Data(in).Insert(ctx)
	if resp.GetError() != nil {
		logx.Errorf("添加日志错误 err:%v", resp.GetError())
		return resp.GetError()
	}
	return nil
}
