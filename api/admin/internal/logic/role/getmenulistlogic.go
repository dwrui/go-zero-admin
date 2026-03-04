// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package role

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList(req *types.GetMenuListReq) (any, error) {
	rpcReq := &system.GetMenuListRequest{
		Pid:    req.Pid,
		UserId: req.UserId,
	}
	rpcResp, err := l.svcCtx.SystemRoleClient.GetMenuList(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	menuData := map[string]interface{}{}
	err = json.Unmarshal([]byte(rpcResp.Data), &menuData)
	if err != nil {
		return nil, err
	}
	return menuData, nil
}
