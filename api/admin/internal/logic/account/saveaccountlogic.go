package account

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAccountLogic {
	return &SaveAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAccountLogic) SaveAccount(req *types.SaveAccountReq) (resp *types.SaveAccountResp, err error) {
	rpcReq := &system.SaveAccountRequest{
		Id:         req.Id,
		Address:    req.Address,
		Avatar:     req.Avatar,
		City:       req.City,
		Company:    req.Company,
		Deptname:   req.Deptname,
		CreateTime: req.CreateTime,
		DeptId:     req.DeptId,
		Email:      req.Email,
		Mobile:     req.Mobile,
		Name:       req.Name,
		Remark:     req.Remark,
		Status:     req.Status,
		Tel:        req.Tel,
		Username:   req.Username,
		Roleid:     req.Roleid,
		Rolename:   req.Rolename,
		Password:   req.Password,
		AccountId:  ga.Uint64(l.ctx.Value("user_id")),
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
	}
	rpcResp, err := l.svcCtx.SystemAccountClient.Save(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &types.SaveAccountResp{
		Id: rpcResp.Id,
	}, nil
}
