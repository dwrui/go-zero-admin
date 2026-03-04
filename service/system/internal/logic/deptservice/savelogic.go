package deptservicelogic

import (
	"context"
	"system/internal/model"
	"system/internal/svc"
	"system/system"
	"time"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveLogic) Save(in *system.SaveDeptRequest) (*system.SaveDeptResponse, error) {
	// todo: add your logic here and delete this line
	dataMap := ga.Map{}
	dataMap["id"] = in.Id
	dataMap["name"] = in.Name
	dataMap["account_id"] = in.AccountId
	dataMap["business_id"] = in.BusinessId
	dataMap["weigh"] = in.Weigh
	dataMap["pid"] = in.Pid
	dataMap["remark"] = in.Remark
	if in.Id == 0 {
		dataMap["create_time"] = time.Now().Format("2006-01-02 15:04:05")
	}
	id, err := model.SaveDept(l.ctx, l.svcCtx, dataMap)
	if err != nil {
		return nil, err
	}
	return &system.SaveDeptResponse{
		Id: id,
	}, nil
}
