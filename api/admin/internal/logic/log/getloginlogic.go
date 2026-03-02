// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package log

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogic {
	return &GetLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginLogic) GetLogin(req *types.GetLoginReq) (any, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := l.svcCtx.SystemLogClient.GetLogin(l.ctx, &system.GetLogListRequest{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Status:     req.Status,
		User:       req.User,
		CreateTime: req.Create_time,
		Ip:         req.Ip,
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
	})
	if err != nil {
		return nil, err
	}

	items := make([]types.GetLogData, 0, len(resp.Items))
	for _, item := range resp.Items {
		items = append(items, types.GetLogData{
			Id:         item.Id,
			Uid:        item.Uid,
			AccountId:  item.AccountId,
			BusinessId: item.BusinessId,
			Type:       item.Type,
			Status:     item.Status,
			Des:        item.Des,
			Ip:         item.Ip,
			Address:    item.Address,
			UserAgent:  item.UserAgent,
			ErrorMsg:   item.ErrorMsg,
			CreateTime: item.CreatedTime,
			User:       item.User,
		})
	}

	return &types.GetLoginResp{
		Items:    items,
		Total:    resp.Total,
		Page:     resp.Page,
		PageSize: resp.PageSize,
	}, nil
}
