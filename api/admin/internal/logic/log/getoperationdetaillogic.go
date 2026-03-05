// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package log

import (
	"context"
	"fmt"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOperationDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOperationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationDetailLogic {
	return &GetOperationDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOperationDetailLogic) GetOperationDetail(req *types.GetOperationDetailReq) (resp *types.GetOperationDetailResp, err error) {
	grpcResp, err := l.svcCtx.SystemLogClient.GetOperationDetail(l.ctx, &system.GetOperationDetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(grpcResp)
	return &types.GetOperationDetailResp{
		Id:             grpcResp.Id,
		User_id:        grpcResp.UserId,
		Account_id:     grpcResp.AccountId,
		Business_id:    grpcResp.BusinessId,
		Operation_type: grpcResp.Type,
		Method:         grpcResp.Method,
		Path:           grpcResp.Path,
		Ip:             grpcResp.Ip,
		Address:        grpcResp.Address,
		Description:    grpcResp.Description,
		Req_headers:    grpcResp.ReqHeaders,
		Req_body:       grpcResp.ReqBody,
		Resp_headers:   grpcResp.RespHeaders,
		Resp_body:      grpcResp.RespBody,
		Status:         grpcResp.Status,
		Duration:       grpcResp.Duration,
		Create_time:    grpcResp.CreateTime,
		Username:       grpcResp.Username,
	}, nil
}
