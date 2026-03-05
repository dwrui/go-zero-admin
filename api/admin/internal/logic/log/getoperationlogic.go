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

type GetOperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationLogic {
	return &GetOperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOperationLogic) GetOperation(req *types.GetOperationReq) (resp *types.GetOperationResp, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	grpcResp, err := l.svcCtx.SystemLogClient.GetOperation(l.ctx, &system.GetOperationRequest{
		UserName:   req.User_name,
		Ip:         req.Ip,
		Status:     req.Status,
		CreateTime: req.Create_time,
		Page:       req.Page,
		Size:       req.Size,
		UserId:     ga.Uint64(l.ctx.Value("user_id")),
		BusinessId: ga.Uint64(l.ctx.Value("business_id")),
	})
	if err != nil {
		return nil, err
	}

	items := make([]types.OperationLogData, 0, len(grpcResp.Items))
	for _, item := range grpcResp.Items {
		items = append(items, types.OperationLogData{
			Id:             item.Id,
			User_id:        item.UserId,
			Account_id:     item.AccountId,
			Business_id:    item.BusinessId,
			Operation_type: item.Type,
			Method:         item.Method,
			Path:           item.Path,
			Ip:             item.Ip,
			Address:        item.Address,
			Description:    item.Description,
			Req_headers:    item.ReqHeaders,
			Req_body:       item.ReqBody,
			Resp_headers:   item.RespHeaders,
			Resp_body:      item.RespBody,
			Status:         item.Status,
			Duration:       item.Duration,
			Create_time:    item.CreateTime,
			User: types.UserInfo{
				Name:     item.User.Name,
				Username: item.User.Username,
				Avatar:   item.User.Avatar,
				Nickname: item.User.Nickname,
			},
		})
	}

	return &types.GetOperationResp{
		Items:    items,
		Page:     grpcResp.Page,
		PageSize: grpcResp.PageSize,
		Total:    grpcResp.Total,
	}, nil
}
