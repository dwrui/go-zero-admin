// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package role

import (
	"context"
	"fmt"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveRoleLogic {
	return &SaveRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveRoleLogic) SaveRole(req *types.SaveRoleReq) (resp *types.SaveRoleResp, err error) {
	fmt.Printf("Request status: %d\n", req.Status) // 打印 status 值
	var menuList []string
	switch v := req.Menu.(type) {
	case string:
		if v == "*" {
			menuList = []string{"*"}
		}
	case []string:
		menuList = v
	case []int64:
		for _, vv := range v {
			menuList = append(menuList, ga.String(vv))
		}
	case []interface{}:
		for _, vv := range v {
			menuList = append(menuList, ga.String(vv))
		}
	}
	rpcReq := &system.SaveRoleRequest{
		Id:         req.Id,
		Name:       req.Name,
		Pid:        req.Pid,
		Status:     req.Status,
		Weigh:      req.Weigh,
		Remark:     req.Remark,
		Rules:      req.Rules,
		Menu:       menuList,
		Btns:       req.Btns,
		DataAccess: req.DataAccess,
		AccountId:  ga.Uint64(l.ctx.Value("user_id")),
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
	}
	rpcResp, err := l.svcCtx.SystemRoleClient.Save(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return &types.SaveRoleResp{
		Id: rpcResp.Id,
	}, nil
}
