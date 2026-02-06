package dashboardservicelogic

import (
	"common/internal/model"
	"context"

	"common/common"
	"common/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuickLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuickLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuickLogic {
	return &GetQuickLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuickLogic) GetQuick(in *common.GetQuickRequest) (*common.GetQuickResponse, error) {
	// todo: add your logic here and delete this line
	list, err := model.GetQuickList(l.ctx, l.svcCtx, in.BusinessId)
	if err != nil {
		return nil, err
	}
	return &common.GetQuickResponse{
		Data: list,
	}, nil
}
