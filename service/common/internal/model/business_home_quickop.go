package model

import (
	"common/common"
	"common/internal/svc"
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

type BusinessAuthQuickop struct {
	Id         uint64 `db:"id"`
	Uid        int64  `db:"uid"`         // 添加人
	BusinessId int64  `db:"business_id"` // 业务主账号id
	IsCommon   int64  `db:"is_common"`   // 公共1=是
	Type       int64  `db:"type"`        // 类型1=外部
	Name       string `db:"name"`        // 快捷名称
	PathUrl    string `db:"path_url"`    // 跳转路径
	Icon       string `db:"icon"`        // 图标
	Weigh      int64  `db:"weigh"`       // 排序
}

func GetQuickList(ctx context.Context, svcCtx *svc.ServiceContext, businessId uint64) (list []*common.GetQuickRow, err error) {
	resp := svcCtx.DB.Model("business_auth_quickop").Where(ga.Map{"business_id": businessId}).Select(ctx, &list)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	return list, nil
}
