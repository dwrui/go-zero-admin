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

func (l *GetDetailLogic) GetDetail(in *configcenter.GetCategoryDetailRequest) (*configcenter.GetCategoryDetailResponse, error) {
	category, err := model.GetCategoryDetail(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	return &configcenter.GetCategoryDetailResponse{
		Data: &configcenter.ConfigCategoryData{
			Id:           ga.Uint64(category.Id),
			CategoryKey:  category.CategoryKey,
			CategoryName: category.CategoryName,
			Description:  category.Description.String,
			SortOrder:    category.SortOrder,
			IsSystem:     category.IsSystem,
			CreateTime:   category.CreateTime.Time.Format("2006-01-02 15:04:05"),
			UpdateTime:   category.UpdateTime.Time.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
