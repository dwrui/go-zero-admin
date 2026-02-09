// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package common

import (
	"admin/grpc-client/common"
	"context"
	"errors"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuLogic) GetMenu(req *types.GetMenuReq) (any, error) {
	userId := ga.Uint64(l.ctx.Value("user_id"))
	if userId == 0 {
		return nil, errors.New("用户不存在")
	}
	resp, err := l.svcCtx.CommonClient.GetMenu(l.ctx, &common.GetMenuRequest{
		RouteId: req.RouteId,
		UserId:  userId,
	})
	if err != nil {
		return nil, err
	}
	menuData := []map[string]interface{}{}
	err = json.Unmarshal([]byte(resp.Data), &menuData)
	if err != nil {
		return nil, err
	}
	return menuData, nil
}
