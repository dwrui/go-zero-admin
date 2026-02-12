package model

import (
	"common/common"
	"common/internal/svc"
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

type CommonAuthQuickop struct {
	Id         uint64 `db:"id"`
	Uid        uint64 `db:"uid"`         // 添加人
	BusinessId uint64 `db:"business_id"` // 业务主账号id
	IsCommon   uint64 `db:"is_common"`   // 公共1=是
	Type       uint64 `db:"type"`        // 类型1=外部
	Name       string `db:"name"`        // 快捷名称
	PathUrl    string `db:"path_url"`    // 跳转路径
	Icon       string `db:"icon"`        // 图标
	Weigh      uint64 `db:"weigh"`       // 排序
}

func GetQuickList(ctx context.Context, svcCtx *svc.ServiceContext, businessId uint64) (list []*common.GetQuickRow, err error) {
	if businessId == 0 {
		businessId = 1
	}
	var commonAuthQuickop []*CommonAuthQuickop
	resp := svcCtx.DB.Model("common_home_quickop").Where(ga.Map{"business_id": businessId}).Select(ctx, &commonAuthQuickop)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	list = ConvertToGetQuickRow(commonAuthQuickop)
	return list, nil
}

// Convert
func ConvertToGetQuickRow(src []*CommonAuthQuickop) []*common.GetQuickRow {
	// 预分配容量（性能优化关键）
	dst := make([]*common.GetQuickRow, 0, len(src))
	for _, a := range src {
		// 直接赋值：结构一致，字段一一对应
		dst = append(dst, &common.GetQuickRow{
			Id:         a.Id,
			Uid:        a.Uid,
			BusinessId: a.BusinessId,
			IsCommon:   a.IsCommon,
			Type:       a.Type,
			Name:       a.Name,
			PathUrl:    a.PathUrl,
			Icon:       a.Icon,
			Weigh:      a.Weigh,
		})
	}
	return dst
}
func SaveQuick(ctx context.Context, svcCtx *svc.ServiceContext, req *common.SaveQuickRequest) (any, error) {
	resp := svcCtx.DB.Model("common_home_quickop").Save(ctx, req)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	return resp.GetLastId(), nil
}
