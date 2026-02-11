package deptservicelogic

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gconv"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

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

func (l *GetListLogic) GetList(in *system.GetDeptListRequest) (*system.GetDeptListResponse, error) {
	list, err := model.GetDeptList(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		newList := make([]map[string]interface{}, 0)
		for _, val := range list {
			newList = append(newList, gconv.Map(val))
		}
		menuLists := ga.GetTreeArray(newList, 0, "")
		var endList []*system.DeptData
		for _, val := range menuLists {
			endList = append(endList, &system.DeptData{
				Id:         ga.Uint64(val["id"]),
				BusinessId: ga.Uint64(val["business_id"]),
				AccountId:  ga.Uint64(val["account_id"]),
				Name:       ga.String(val["name"]),
				Pid:        ga.Uint64(val["pid"]),
				Weigh:      ga.Uint64(val["weigh"]),
				Status:     ga.Uint64(val["status"]),
				Remark:     ga.String(val["remark"]),
				CreateTime: ga.String(val["create_time"]),
				Spacer:     ga.String(val["spacer"]),
				Children:   val["children"].([]*system.DeptData),
			})
		}
		return &system.GetDeptListResponse{
			Data: endList,
		}, nil
	}
	return &system.GetDeptListResponse{}, nil
}
