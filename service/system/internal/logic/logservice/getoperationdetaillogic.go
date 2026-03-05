package logservicelogic

import (
	"context"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOperationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOperationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationDetailLogic {
	return &GetOperationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOperationDetailLogic) GetOperationDetail(in *system.GetOperationDetailRequest) (*system.GetOperationDetailResponse, error) {
	// todo: add your logic here and delete this line
	resp, err := model.GetOperatuinDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	return &system.GetOperationDetailResponse{
		Id:          resp.Id,
		UserId:      resp.UserId,
		AccountId:   resp.AccountId,
		BusinessId:  resp.BusinessId,
		Type:        resp.Type,
		Method:      resp.Method,
		Path:        resp.Path,
		Ip:          resp.IP,
		Address:     resp.Address,
		Description: resp.Description,
		ReqHeaders:  resp.ReqHeaders,
		ReqBody:     resp.ReqBody,
		RespHeaders: resp.RespHeaders,
		RespBody:    resp.RespBody,
		Status:      resp.Status,
		Duration:    resp.Duration,
		CreateTime:  resp.CreateTime.Time.Format("2006-01-02 15:04:05"),
		Username:    resp.UserName.String,
	}, nil
}
