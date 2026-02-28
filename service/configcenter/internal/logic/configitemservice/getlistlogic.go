package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"encoding/json"
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

func (l *GetListLogic) GetList(in *configcenter.GetConfigListRequest) (*configcenter.GetConfigListResponse, error) {
	list, err := model.GetConfigList(l.ctx, l.svcCtx, in)
	if err != nil {
		return nil, err
	}

	items, ok := list["items"].([]*model.CommonConfigItem)
	if !ok {
		return nil, errors.New("items 类型错误")
	}
	configItemList := make([]*configcenter.ConfigItemData, 0)
	for _, item := range items {
		configItemList = append(configItemList, convertToConfigItemData(item))
	}

	return &configcenter.GetConfigListResponse{
		Items:    configItemList,
		Total:    ga.Uint64(list["total"]),
		Page:     ga.Uint64(in.Page),
		PageSize: ga.Uint64(in.PageSize),
	}, nil
}

func convertToConfigItemData(item *model.CommonConfigItem) *configcenter.ConfigItemData {
	data := &configcenter.ConfigItemData{
		Id:             ga.Uint64(item.Id),
		CategoryKey:    item.CategoryKey,
		ConfigKey:      item.ConfigKey,
		ConfigName:     item.ConfigName,
		ConfigType:     item.ConfigType,
		ConfigValue:    item.ConfigValue.String,
		DefaultValue:   item.DefaultValue.String,
		Description:    item.Description.String,
		ValidationRule: item.ValidationRule.String,
		Placeholder:    item.Placeholder.String,
		IsRequired:     item.IsRequired,
		IsSecret:       item.IsSecret,
		SortOrder:      item.SortOrder,
		Status:         item.Status,
		CreatedTime:    item.CreateTime.Time.Format("2006-01-02 15:04:05"),
		UpdatedTime:    item.UpdateTime.Time.Format("2006-01-02 15:04:05"),
	}

	// 解析options JSON
	if item.Options.Valid && item.Options.String != "" {
		var options []*configcenter.ConfigOptionItem
		if err := json.Unmarshal([]byte(item.Options.String), &options); err != nil {
			logx.Errorf("解析options失败: %v", err)
			data.Options = []*configcenter.ConfigOptionItem{}
		} else {
			data.Options = options
		}
	} else {
		data.Options = []*configcenter.ConfigOptionItem{}
	}

	return data
}
