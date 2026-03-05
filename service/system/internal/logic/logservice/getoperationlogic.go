package logservicelogic

import (
	"context"
	"net"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOperationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOperationLogic {
	return &GetOperationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOperation 获取操作日志列表
func (l *GetOperationLogic) GetOperation(in *system.GetOperationRequest) (*system.GetOperationResponse, error) {
	whereMap := gmap.New()
	if in.UserId != 1 {
		whereMap.Set("log.business_id", in.BusinessId)
	}
	if in.Page == 0 {
		whereMap.Set("page", 1)
	}
	if in.Size == 0 {
		whereMap.Set("pageSize", 10)
	}
	if in.UserName != "" {
		userids := l.svcCtx.DB.Model("admin_account").Where("name like ?", "%"+ga.String(in.UserName)+"%").Column(l.ctx, "id", &[]uint64{})
		whereMap.Set("log.uid IN(?)", userids.GetData().([]string))
	}
	if in.Ip != "" {
		address := net.ParseIP(ga.String(in.Ip))
		if address == nil {
			whereMap.Set("log.address like ?", "%"+ga.String(address)+"%")
		} else {
			whereMap.Set("log.ip", in.Ip)
		}
	}
	if in.Status == 0 || in.Status == 1 {
		whereMap.Set("log.status", in.Status)
	}
	if in.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(ga.String(in.CreateTime), ",")
		whereMap.Set("log.create_time between ? and ?", ga.Slice{datetime_arr[0], datetime_arr[1]})
	}
	resp, err := model.GetOperationLogList(l.ctx, l.svcCtx, whereMap, in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	//整合数据
	items := resp["items"].([]*model.LogOperationWithUserInfo)
	returnData := []*system.OperationLogData{}
	for _, i := range items {
		returnData = append(returnData, &system.OperationLogData{
			Id:          i.Id,
			BusinessId:  i.BusinessId,
			UserId:      i.UserId,
			AccountId:   i.AccountId,
			Type:        i.Type,
			Method:      i.Method,
			Status:      i.Status,
			Address:     i.Address,
			Description: i.Description,
			Path:        i.Path,
			Ip:          i.IP,
			ReqHeaders:  i.ReqHeaders,
			ReqBody:     i.ReqBody,
			RespHeaders: i.RespHeaders,
			RespBody:    i.RespBody,
			Duration:    i.Duration,
			User: &system.UserInfo{
				Name:     i.Name.String,
				Username: i.UserName.String,
				Avatar:   i.Avatar.String,
				Nickname: i.UserNickname.String,
			},
			CreateTime: i.CreateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return &system.GetOperationResponse{
		Items:    returnData,
		Page:     in.Page,
		PageSize: in.Size,
		Total:    ga.Uint64(resp["total"]),
	}, nil
}
