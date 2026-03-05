package account

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccountRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountRoleLogic {
	return &GetAccountRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountRoleLogic) GetAccountRole(req *types.GetAccountRoleReq) (resp *types.GetAccountRoleResp, err error) {
	rpcResp, err := l.svcCtx.SystemAccountClient.GetRole(l.ctx, &system.GetAccountRoleRequest{
		UserId: ga.Uint64(l.ctx.Value("user_id")),
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.RoleData, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		children := make([]types.RoleData, 0)
		for _, child := range item.Children {
			children = append(children, types.RoleData{
				AccountId:  child.AccountId,
				Btns:       child.Btns,
				BusinessId: child.BusinessId,
				CreateTime: child.CreateTime,
				DataAccess: child.DataAccess,
				Id:         child.Id,
				Menu:       child.Menu,
				Name:       child.Name,
				Pid:        child.Pid,
				Remark:     child.Remark,
				Rules:      child.Rules,
				Spacer:     child.Spacer,
				Status:     child.Status,
				Weigh:      child.Weigh,
				Children:   []types.RoleData{},
			})
		}
		list = append(list, types.RoleData{
			AccountId:  item.AccountId,
			Btns:       item.Btns,
			BusinessId: item.BusinessId,
			CreateTime: item.CreateTime,
			DataAccess: item.DataAccess,
			Id:         item.Id,
			Menu:       item.Menu,
			Name:       item.Name,
			Pid:        item.Pid,
			Remark:     item.Remark,
			Rules:      item.Rules,
			Spacer:     item.Spacer,
			Status:     item.Status,
			Weigh:      item.Weigh,
			Children:   children,
		})
	}

	return &types.GetAccountRoleResp{
		List: list,
	}, nil
}
