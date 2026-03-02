package deptservicelogic

import (
	"context"
	"system/internal/model"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"

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
			newList = append(newList, ga.Map{
				"id":          val.Id,
				"pid":         val.Pid,
				"name":        val.Name,
				"account_id":  val.AccountId,
				"weigh":       val.Weigh,
				"status":      val.Status,
				"remark":      val.Remark,
				"business_id": val.BusinessId,
				"create_time": val.CreateTime.Time.Format("2006-01-02 15:04:05"),
			})
		}
		menuLists := ga.GetMenuChildrenArray(newList, 0, "pid")

		var convertToDeptData func(menuList map[string]interface{}) *system.DeptParentData
		convertToDeptData = func(menuList map[string]interface{}) *system.DeptParentData {
			// 处理children字段，从gvar.Var中提取实际数据
			var children []*system.DeptParentData
			if childrenData := menuList["children"]; childrenData != nil {
				// 直接使用ga.Interfaces函数提取数据
				childrenInterfaces := ga.Interfaces(childrenData)
				for _, childInterface := range childrenInterfaces {
					if childMap, ok := childInterface.(map[string]interface{}); ok {
						children = append(children, convertToDeptData(childMap))
					}
				}
			}

			return &system.DeptParentData{
				Id:       ga.Uint64(menuList["id"]),
				Name:     ga.String(menuList["name"]),
				Pid:      ga.Uint64(menuList["pid"]),
				Children: children,
			}
		}
		var endList []*system.DeptParentData
		for _, val := range menuLists {
			endList = append(endList, convertToDeptData(val))
		}
		return &system.GetDeptParentResponse{
			Data: endList,
		}, nil
	}
	return &system.GetDeptParentResponse{}, nil
}
