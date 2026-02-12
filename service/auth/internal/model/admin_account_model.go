package model

import (
	"auth/internal/svc"
	"context"
	"database/sql"
	"errors"
	"time"
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

/**
 * 获取用户信息
 * @param ctx 上下文
 * @param svg 服务上下文
 * @param id 用户id
 * @param field 字段
 */
func GetUserInfo(ctx context.Context, svg *svc.ServiceContext, id uint64, field string) (AdminAccountModel, error) {
	userModel := AdminAccountModel{}
	resp := svg.DB.Model("admin_account").Fields(field).Where("id = ?", id).Find(ctx, &userModel)
	if resp.GetError() != nil || resp.IsEmpty() {
		return AdminAccountModel{}, errors.New("用户不存在")
	}
	return userModel, nil
}
