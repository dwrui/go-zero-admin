package logservicelogic

import (
	"apilog/internal/model"
	"context"
	"time"

	"apilog/apilog"
	"apilog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOperationLogLogic {
	return &AddOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddOperationLogLogic) AddOperationLog(in *apilog.OperationLogRequest) (*apilog.OperationLogResponse, error) {
	// todo: add your logic here and delete this line
	des := model.GetDes(l.ctx, l.svcCtx, in.Path)
	in.Description = des
	in.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	err := model.AddOperationLog(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	return &apilog.OperationLogResponse{}, nil
}
