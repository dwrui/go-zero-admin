package roleservicelogic

import (
	"context"
	"errors"
	"system/internal/model"

	"system/internal/svc"
	"system/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuListLogic) GetMenuList(in *system.GetMenuListRequest) (*system.GetMenuListResponse, error) {
	// todo: add your logic here and delete this line
	resp, err := model.GetRoleMenuList(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	dataJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	return &system.GetMenuListResponse{
		Data: ga.String(dataJson),
	}, nil
}
