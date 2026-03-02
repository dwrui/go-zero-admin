package deptservicelogic

import (
	"context"
	"system/internal/model"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"

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
	whereMap := gmap.New()
	if in.BusinessId != 0 && in.BusinessId != 1 {
		whereMap.Set("business_id =?", in.BusinessId)
	}
	if in.Name != "" {
		whereMap.Set("name like ?", "%"+in.Name+"%")
	}
	if in.Status == 0 || in.Status == 1 {
		whereMap.Set("status =?", in.Status)
	}
	if in.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(ga.String(in.CreateTime), ",")
		whereMap.Set("create_time between ? and ?", ga.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	list, err := model.GetDeptList(l.ctx, l.svcCtx, whereMap)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		newList := make([]map[string]interface{}, 0)
		for _, val := range list {
			newList = append(newList, ga.Map{
				"id":          val.Id,
				"business_id": val.BusinessId,
				"account_id":  val.AccountId,
				"name":        val.Name,
				"pid":         val.Pid,
				"weigh":       val.Weigh,
				"status":      val.Status,
				"remark":      val.Remark,
				"create_time": val.CreateTime.Time.Format("2006-01-02 15:04:05"),
			})
		}
		menuLists := ga.GetTreeArray(newList, 0, "")
		// 递归函数，将map转换为DeptData结构
		var convertToDeptData func(menuList map[string]interface{}) *system.DeptData
		convertToDeptData = func(menuList map[string]interface{}) *system.DeptData {
			// 处理children字段，从gvar.Var中提取实际数据
			var children []*system.DeptData
			if childrenData := menuList["children"]; childrenData != nil {
				// 直接使用ga.Interfaces函数提取数据
				childrenInterfaces := ga.Interfaces(childrenData)
				for _, childInterface := range childrenInterfaces {
					if childMap, ok := childInterface.(map[string]interface{}); ok {
						children = append(children, convertToDeptData(childMap))
					}
				}
			}

			return &system.DeptData{
				Id:         ga.Uint64(menuList["id"]),
				BusinessId: ga.Uint64(menuList["business_id"]),
				AccountId:  ga.Uint64(menuList["account_id"]),
				Name:       ga.String(menuList["name"]),
				Pid:        ga.Uint64(menuList["pid"]),
				Weigh:      ga.Uint64(menuList["weigh"]),
				Status:     ga.Uint64(menuList["status"]),
				Remark:     ga.String(menuList["remark"]),
				CreateTime: ga.String(menuList["create_time"]),
				Children:   children,
			}
		}
		var endList []*system.DeptData
		for _, val := range menuLists {
			endList = append(endList, convertToDeptData(val))
		}
		return &system.GetDeptListResponse{
			Data: endList,
		}, nil
	}
	return &system.GetDeptListResponse{}, nil
}
