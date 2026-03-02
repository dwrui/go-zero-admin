// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package dept

import (
	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"
	"context"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeptParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeptParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeptParentLogic {
	return &GetDeptParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeptParentLogic) GetDeptParent(req *types.GetDeptParentReq) (resp *types.GetDeptParentResp, err error) {
	rpcReq := &system.GetDeptParentRequest{
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
	}
	rpcResp, err := l.svcCtx.SystemDeptClient.GetParent(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	data := make([]types.DeptParentData, 0)
	for _, item := range rpcResp.Data {
		children := make([]types.DeptParentData, 0)
		for _, child := range item.Children {
			children = append(children, types.DeptParentData{
				Id:   child.Id,
				Name: child.Name,
				Pid:  child.Pid,
			})
		}
		data = append(data, types.DeptParentData{
			Id:       item.Id,
			Name:     item.Name,
			Pid:      item.Pid,
			Children: children,
		})
	}
	return &types.GetDeptParentResp{
		Data: data,
	}, nil
}
