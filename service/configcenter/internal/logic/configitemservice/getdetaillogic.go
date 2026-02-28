package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDetailLogic) GetDetail(in *configcenter.GetConfigDetailRequest) (*configcenter.GetConfigDetailResponse, error) {
	item, err := model.GetConfigDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("配置项不存在")
	}

	return &configcenter.GetConfigDetailResponse{
		Data: convertToConfigItemData(item),
	}, nil
}
