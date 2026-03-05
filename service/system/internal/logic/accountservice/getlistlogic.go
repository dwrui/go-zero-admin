package accountservicelogic

import (
	"context"
	"system/internal/model"
	"system/internal/svc"
	"system/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *system.GetAccountListRequest) (*system.GetAccountListResponse, error) {
	whereMap := gmap.New()
	if in.UserId != 1 {
		whereMap.Set("business_id", in.BusinessId)
	}
	account_id, filter := model.GetDataAuthor(l.ctx, l.svcCtx, in.UserId, "")
	if filter {
		whereMap.Set("id IN(?)", account_id) //in 查询
	}

	if in.Name != "" {
		whereMap.Set("name like ?", "%"+in.Name+"%")
	}
	if in.Status == 0 || in.Status == 1 {
		whereMap.Set("status = ?", in.Status)
	}
	if in.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(in.CreateTime, ",")
		whereMap.Set("create_time between ? and ?", ga.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	resp, err := model.GetAccountList(l.ctx, l.svcCtx, whereMap, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	list := resp["list"].([]map[string]interface{})
	var accountList []*system.AccountData
	for _, list_one := range list {
		accountList = append(accountList, &system.AccountData{
			Id:         ga.Uint64(list_one["id"]),
			Status:     ga.Uint64(list_one["status"]),
			Name:       ga.String(list_one["name"]),
			Username:   ga.String(list_one["username"]),
			Avatar:     ga.String(list_one["avatar"]),
			Tel:        ga.String(list_one["tel"]),
			Mobile:     ga.String(list_one["mobile"]),
			Email:      ga.String(list_one["email"]),
			DeptId:     ga.Uint64(list_one["dept_id"]),
			Remark:     ga.String(list_one["remark"]),
			City:       ga.String(list_one["city"]),
			Address:    ga.String(list_one["address"]),
			Company:    ga.String(list_one["company"]),
			CreateTime: ga.String(list_one["create_time"]),
			Deptname:   ga.String(list_one["dept_name"]),
			Roleid:     list_one["role_id"].([]uint64),
			Rolename:   list_one["role_name"].([]string),
		})
	}
	return &system.GetAccountListResponse{
		Items:    accountList,
		Page:     in.Page,
		PageSize: in.PageSize,
		Total:    ga.Uint64(resp["total"]),
	}, nil
}
