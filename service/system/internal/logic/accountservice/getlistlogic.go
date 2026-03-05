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
		whereMap.Set("c.id IN(?)", account_id) //in 查询
	}

	if in.Name != "" {
		whereMap.Set("c.name like ?", "%"+in.Name+"%")
	}
	if in.Status == 0 || in.Status == 1 {
		whereMap.Set("c.status = ?", in.Status)
	}
	if in.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(in.CreateTime, ",")
		whereMap.Set("c.create_time between ? and ?", ga.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	resp, err := model.GetAccountList(l.ctx, l.svcCtx, whereMap, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	list := resp["list"].([]*model.AdminAccountModelResponse)
	var accountList []*system.AccountData
	for _, list_one := range list {
		accountList = append(accountList, &system.AccountData{
			Id:         list_one.Id,
			Status:     ga.Uint64(list_one.Status),
			Name:       list_one.Name,
			Username:   list_one.Username,
			Avatar:     list_one.Avatar,
			Tel:        list_one.Tel,
			Mobile:     list_one.Mobile,
			Email:      list_one.Email,
			DeptId:     ga.Uint64(list_one.DeptId),
			Remark:     list_one.Remark,
			City:       list_one.City,
			Address:    list_one.Address,
			Company:    list_one.Company,
			CreateTime: list_one.CreateTime.Time.Format("2006-01-02 15:04:05"),
			Deptname:   list_one.DeptName,
			Roleid:     list_one.RoleId,
			Rolename:   list_one.RoleName,
		})
	}
	return &system.GetAccountListResponse{
		Items:    accountList,
		Page:     in.Page,
		PageSize: in.PageSize,
		Total:    ga.Uint64(resp["total"]),
	}, nil
}
