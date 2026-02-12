package commonservicelogic

import (
	"common/internal/model"
	"context"
	"errors"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gvar"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"

	"common/common"
	"common/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuLogic) GetMenu(in *common.GetMenuRequest) (*common.GetMenuResponse, error) {
	if in.UserId == 0 {
		if in.RouteId == 0 {
			return &common.GetMenuResponse{}, errors.New("route_id 不能为空")
		}
		var routeInfo []model.AdminAuthRule
		nemu_list := l.svcCtx.DB.Model("admin_auth_rule").Where("id", in.RouteId).Order("weigh", "asc").Select(l.ctx, &routeInfo)
		if nemu_list.IsEmpty() {
			return &common.GetMenuResponse{}, errors.New("路由不存在")
		}
		rulemenu := model.GetMenuArray(l.ctx, l.svcCtx, routeInfo, 0, make([]interface{}, 0))
		// 将数据转换为JSON字符串
		jsonData, err := json.Marshal(rulemenu)
		if err != nil {
			return &common.GetMenuResponse{}, err
		}
		return &common.GetMenuResponse{
			Data: ga.String(jsonData),
		}, nil
	}
	//获取用户权限菜单
	role_id := l.svcCtx.DB.Model("admin_auth_role_access").Where("uid", in.UserId).Column(l.ctx, "role_id")
	if role_id.GetError() != nil {
		return &common.GetMenuResponse{}, role_id.GetError()
	}
	if role_id.IsEmpty() {
		return &common.GetMenuResponse{}, errors.New("您没有管理后台权限，请联系管理员授权")
	}
	menu_ids := l.svcCtx.DB.Model("admin_auth_role").WhereIn("id", ga.FormatColumnData(role_id.GetData())).Column(l.ctx, "rules")
	if menu_ids.GetError() != nil {
		return &common.GetMenuResponse{}, menu_ids.GetError()
	}
	//获取超级角色
	super_role := l.svcCtx.DB.Model("admin_auth_role").WhereIn("id", ga.FormatColumnData(role_id.GetData())).Where("rules", "*").Value(l.ctx, "id")
	RMDB := l.svcCtx.DB.Model("admin_auth_rule")
	var roles []interface{}
	if super_role == nil { //不是超级权限-过滤菜单权限
		// 获取菜单ID数据并处理格式
		menuData := menu_ids.GetData()
		if menuData != nil {
			if dataSlice, ok := menuData.([]interface{}); ok && len(dataSlice) > 0 {
				// 创建 []*gvar.Var 数组
				varList := make([]*gvar.Var, len(dataSlice))
				for i, v := range dataSlice {
					varList[i] = ga.VarNew(v)
				}
				getmenus := ga.ArrayMerge(varList)
				RMDB = RMDB.WhereIn("id", getmenus)
				roles = getmenus
			}
		}
	} else {
		roles = make([]interface{}, 0)
	}
	var menuInfo []model.AdminAuthRule
	nemu_list := RMDB.Where("status", 0).WhereIn("type", ga.Slice{0, 1}).OrderBy("weigh").Select(l.ctx, &menuInfo)
	if nemu_list.GetError() != nil {
		return &common.GetMenuResponse{}, nemu_list.GetError()
	}

	rulemenu := model.GetMenuArray(l.ctx, l.svcCtx, menuInfo, 0, roles)
	// 将数据转换为JSON字符串
	jsonData, err := json.Marshal(rulemenu)
	if err != nil {
		return &common.GetMenuResponse{}, err
	}
	return &common.GetMenuResponse{
		Data: ga.String(jsonData),
	}, nil
}
