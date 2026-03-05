package account

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountListLogic {
	return &GetAccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountListLogic) GetAccountList(req *types.GetAccountListReq) (resp *types.GetAccountListResp, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	rpcReq := &system.GetAccountListRequest{
		UserId:     ga.Uint64(l.ctx.Value("user_id")),
		Name:       req.Name,
		Status:     req.Status,
		CreateTime: req.CreateTime,
		Page:       req.Page,
		PageSize:   req.PageSize,
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
	}
	rpcResp, err := l.svcCtx.SystemAccountClient.GetList(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	items := make([]types.AccountData, 0, len(rpcResp.Items))
	for _, item := range rpcResp.Items {
		items = append(items, types.AccountData{
			Id:         item.Id,
			Address:    item.Address,
			Avatar:     item.Avatar,
			City:       item.City,
			Company:    item.Company,
			Deptname:   item.Deptname,
			CreateTime: item.CreateTime,
			DeptId:     item.DeptId,
			Email:      item.Email,
			Mobile:     item.Mobile,
			Name:       item.Name,
			Remark:     item.Remark,
			Status:     item.Status,
			Tel:        item.Tel,
			Username:   item.Username,
			Roleid:     item.Roleid,
			Rolename:   item.Rolename,
		})
	}

	return &types.GetAccountListResp{
		Items:    items,
		Page:     rpcResp.Page,
		PageSize: rpcResp.PageSize,
		Total:    rpcResp.Total,
	}, nil
}