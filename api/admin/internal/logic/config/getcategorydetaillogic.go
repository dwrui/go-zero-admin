// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"admin/grpc-client/configcenter"
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryDetailLogic {
	return &GetCategoryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryDetailLogic) GetCategoryDetail(req *types.GetCategoryDetailReq) (resp *types.ConfigCategoryData, err error) {
	rpcResp, err := l.svcCtx.ConfigCenterClient.GetDetail(l.ctx, &configcenter.GetCategoryDetailRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("获取配置分类详情失败: %v", err)
		return nil, err
	}

	return &types.ConfigCategoryData{
		Id:           rpcResp.Data.Id,
		CategoryKey:  rpcResp.Data.CategoryKey,
		CategoryName: rpcResp.Data.CategoryName,
		Description:  rpcResp.Data.Description,
		SortOrder:    rpcResp.Data.SortOrder,
		IsSystem:     rpcResp.Data.IsSystem,
		CreateTime:   rpcResp.Data.CreateTime,
		UpdateTime:   rpcResp.Data.UpdateTime,
	}, nil
}
