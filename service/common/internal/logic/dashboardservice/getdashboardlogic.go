package dashboardservicelogic

import (
	"context"

	"common/common"
	"common/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDashboardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDashboardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDashboardLogic {
	return &GetDashboardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDashboardLogic) GetDashboard(in *common.GetQuickRequest) (*common.GetQuickResponse, error) {
	// todo: add your logic here and delete this line

	return &common.GetQuickResponse{}, nil
}
