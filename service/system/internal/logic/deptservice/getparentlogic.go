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

type GetParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParentLogic {
	return &GetParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetParentLogic) GetParent(in *system.GetDeptParentRequest) (*system.GetDeptParentResponse, error) {
	parentList, err := model.GetDeptParent(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	if len(parentList) > 0 {
		newList := make([]map[string]interface{}, 0)
		for _, val := range parentList {
			newList = append(newList, gconv.Map(val))
		}
		menuLists := ga.GetMenuChildrenArray(newList, 0, "pid")
		var list []*system.DeptParentData
		for _, v := range menuLists {
			list = append(list, &system.DeptParentData{
				Id:       ga.Uint64(v["id"]),
				Name:     ga.String(v["name"]),
				Pid:      ga.Uint64(v["pid"]),
				Children: v["children"].([]*system.DeptParentData),
			})
		}
		return &system.GetDeptParentResponse{
			Data: list,
		}, nil
	}
	return &system.GetDeptParentResponse{}, nil
}
