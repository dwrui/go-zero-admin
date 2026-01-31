package model

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"user/internal/svc"
)

var _ GaBusinessAccountModel = (*customGaBusinessAccountModel)(nil)

type (
	// GaBusinessAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGaBusinessAccountModel.
	GaBusinessAccountModel interface {
		gaBusinessAccountModel
		withSession(session sqlx.Session) GaBusinessAccountModel
	}

	customGaBusinessAccountModel struct {
		*defaultGaBusinessAccountModel
	}
)

// NewGaBusinessAccountModel returns a model for the database table.
func NewGaBusinessAccountModel(conn sqlx.SqlConn) GaBusinessAccountModel {
	return &customGaBusinessAccountModel{
		defaultGaBusinessAccountModel: newGaBusinessAccountModel(conn),
	}
}

func (m *customGaBusinessAccountModel) withSession(session sqlx.Session) GaBusinessAccountModel {
	return NewGaBusinessAccountModel(sqlx.NewSqlConnFromSession(session))
}

/**
 * 锁定账户
 * @param ctx
 * @param svg
 * @param id
 */
func LockAccount(ctx context.Context, svg *svc.ServiceContext, id uint64) error {
	LockTime := time.Now().Add(30 * time.Minute)
	res := svg.DB.Model("business_account").Where("id = ?", id).Update(ctx, ga.Map{"login_attempts": 0, "lock_time": LockTime})
	if res.GetError() != nil {
		return res.GetError()
	}
	return nil
}
