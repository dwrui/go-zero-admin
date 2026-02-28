package configcategoryservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
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

func (l *GetListLogic) GetList(in *configcenter.GetCategoryListRequest) (*configcenter.GetCategoryListResponse, error) {
	list, err := model.GetCategoryList(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}

	items, ok := list["items"].([]*model.CommonConfigCategory)
	if !ok {
		return nil, errors.New("items 类型错误")
	}
	commonConfigList := make([]*configcenter.ConfigCategoryData, 0)
	for _, item := range items {
		commonConfigList = append(commonConfigList, &configcenter.ConfigCategoryData{
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

	return &configcenter.GetCategoryListResponse{
		Items:    commonConfigList,
		Total:    ga.Uint64(list["total"]),
		Page:     ga.Uint64(in.Page),
		PageSize: ga.Uint64(in.PageSize),
	}, nil
}
