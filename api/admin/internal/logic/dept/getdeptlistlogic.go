// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package dept

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeptListLogic {
	return &GetDeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeptListLogic) GetDeptList(req *types.GetDeptListReq) (resp *types.GetDeptListResp, err error) {
	rpcReq := &system.GetDeptListRequest{
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
		Name:       req.Name,
		Status:     req.Status,
		CreateTime: req.CreateTime,
	}
	rpcResp, err := l.svcCtx.SystemDeptClient.GetList(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	data := make([]types.DeptData, 0)
	for _, item := range rpcResp.Data {
		children := make([]types.DeptData, 0)
		for _, child := range item.Children {
			children = append(children, types.DeptData{
				Id:         child.Id,
				BusinessId: child.BusinessId,
				AccountId:  child.AccountId,
				Name:       child.Name,
				Pid:        child.Pid,
				Weigh:      child.Weigh,
				Status:     child.Status,
				Remark:     child.Remark,
				CreateTime: child.CreateTime,
				Spacer:     child.Spacer,
				Children:   []types.DeptData{},
			})
		}
		data = append(data, types.DeptData{
			Id:         item.Id,
			BusinessId: item.BusinessId,
			AccountId:  item.AccountId,
			Name:       item.Name,
			Pid:        item.Pid,
			Weigh:      item.Weigh,
			Status:     item.Status,
			Remark:     item.Remark,
			CreateTime: item.CreateTime,
			Spacer:     item.Spacer,
			Children:   children,
		})
	}
	return &types.GetDeptListResp{
		Data: data,
	}, nil
}
