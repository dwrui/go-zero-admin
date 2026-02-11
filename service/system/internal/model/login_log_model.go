package model

import (
	"time"
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
