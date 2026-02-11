package dashboardservicelogic

import (
	"common/internal/model"
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"

	"common/common"
	"common/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveQuickLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveQuickLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveQuickLogic {
	return &SaveQuickLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveQuickLogic) SaveQuick(in *common.SaveQuickRequest) (*common.SaveQuickResponse, error) {
	// todo: add your logic here and delete this line
	resp, err := model.SaveQuick(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}
	return &common.SaveQuickResponse{
		Id: ga.Uint64(resp),
	}, nil
}
