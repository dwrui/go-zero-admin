package configcategoryservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllLogic) GetAll(in *configcenter.GetCategoryAllRequest) (*configcenter.GetCategoryAllResponse, error) {
	list, err := model.GetCategoryAll(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}

	var items []*configcenter.ConfigCategoryData
	for _, item := range list {
		items = append(items, &configcenter.ConfigCategoryData{
			Id:           ga.Uint64(item.Id),
			CategoryKey:  item.CategoryKey,
			CategoryName: item.CategoryName,
			Description:  item.Description.String,
			SortOrder:    item.SortOrder,
			IsSystem:     item.IsSystem,
			CreateTime:   item.CreateTime.Time.Format("2006-01-02 15:04:05"),
			UpdateTime:   item.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return &configcenter.GetCategoryAllResponse{
		Items: items,
	}, nil
}
