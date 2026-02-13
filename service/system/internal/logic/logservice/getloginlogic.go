package logservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogic {
	return &GetLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLoginLogic) GetLogin(in *system.GetLogListRequest) (*system.GetLogListResponse, error) {
	// todo: add your logic here and delete this line
	list, err := model.GetLoginLogList(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	items, ok := list["items"].([]*model.LoginLogModel)
	if !ok {
		return nil, errors.New("数据类型错误")
	}
	logList := make([]*system.GetLogData, 0)
	for _, v := range items {
		logList = append(logList, &system.GetLogData{
			Id:          v.Id,
			Uid:         v.Uid,
			AccountId:   v.AccountId,
			BusinessId:  v.BusinessId,
			Type:        v.Type,
			Status:      ga.Uint64(v.Status),
			Des:         v.Des,
			Ip:          v.Ip,
			Address:     v.Address,
			UserAgent:   v.UserAgent,
			User:        v.User,
			CreatedTime: v.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	fmt.Println(logList)
	return &system.GetLogListResponse{
		Items:    logList,
		Page:     ga.Uint64(list["page"]),
		PageSize: ga.Uint64(list["page_size"]),
		Total:    ga.Uint64(list["total"]),
	}, nil
}
