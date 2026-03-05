package accountservicelogic

import (
	"context"
	"fmt"
	"system/internal/model"
	"time"

	"system/internal/svc"
	"system/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/grand"
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

func (l *SaveLogic) Save(in *system.SaveAccountRequest) (*system.SaveAccountResponse, error) {
	dataMap := ga.Map{
		"id":       in.Id,
		"address":  in.Address,
		"avatar":   in.Avatar,
		"city":     in.City,
		"company":  in.Company,
		"dept_id":  in.DeptId,
		"email":    in.Email,
		"mobile":   in.Mobile,
		"name":     in.Name,
		"remark":   in.Remark,
		"status":   in.Status,
		"tel":      in.Tel,
		"username": in.Username,
		"password": in.Password,
	}
	if dataMap["password"] != "" {
		salt := grand.Str("123456789", 6)
		mdpass := fmt.Sprintf("%v%v", dataMap["password"], salt)
		dataMap["password"] = ga.Md5(mdpass)
		dataMap["salt"] = salt
	} else {
		salt := grand.Str("123456789", 6)
		mdpass := fmt.Sprintf("%v%v", "123456", salt)
		dataMap["password"] = ga.Md5(mdpass)
		dataMap["salt"] = salt
	}
	if dataMap["avatar"] == "" {
		dataMap["avatar"] = "/static/unknown.png"
	}
	if in.Id == 0 {
		dataMap["account_id"] = in.AccountId
		dataMap["business_id"] = in.BusinessId
		dataMap["create_time"] = time.Now().Format("2006-01-02 15:04:05")
	} else {
		dataMap["update_time"] = time.Now().Format("2006-01-02 15:04:05")
	}
	id, err := model.SaveAccount(l.ctx, l.svcCtx, dataMap)
	if err != nil {
		return nil, err
	}
	model.AppRoleAccess(l.ctx, l.svcCtx, in.Roleid, id)
	return &system.SaveAccountResponse{
		Id: id,
	}, nil
}
