package model

import (
	"configcenter/configcenter"
	"configcenter/internal/svc"
	"context"
	"database/sql"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
)

type CommonConfigItem struct {
	Id             int64          `json:"id" db:"id"`
	CategoryKey    string         `json:"category_key" db:"category_key"`
	ConfigKey      string         `json:"config_key" db:"config_key"`
	ConfigName     string         `json:"config_name" db:"config_name"`
	ConfigType     string         `json:"config_type" db:"config_type"`
	ConfigValue    sql.NullString `json:"config_value" db:"config_value"`
	DefaultValue   sql.NullString `json:"default_value" db:"default_value"`
	Description    sql.NullString `json:"description" db:"description"`
	Options        sql.NullString `json:"options" db:"options"`
	ValidationRule sql.NullString `json:"validation_rule" db:"validation_rule"`
	Placeholder    sql.NullString `json:"placeholder" db:"placeholder"`
	IsRequired     int32          `json:"is_required" db:"is_required"`
	IsSecret       int32          `json:"is_secret" db:"is_secret"`
	SortOrder      int32          `json:"sort_order" db:"sort_order"`
	Status         int32          `json:"status" db:"status"`
	CreateTime     sql.NullTime   `json:"create_time" db:"create_time"`
	UpdateTime     sql.NullTime   `json:"update_time" db:"update_time"`
}

func GetConfigList(ctx context.Context, svcCtx *svc.ServiceContext, req *configcenter.GetConfigListRequest) (ga.Map, error) {
	whereMap := gmap.New()
	if req.CategoryKey != "" {
		whereMap.Set("category_key = ?", req.CategoryKey)
	}
	if req.ConfigKey != "" {
		whereMap.Set("config_key like ?", "%"+req.ConfigKey+"%")
	}
	if req.ConfigName != "" {
		whereMap.Set("config_name like ?", "%"+req.ConfigName+"%")
	}
	if req.Status != 0 {
		whereMap.Set("status = ?", req.Status)
	}

	var list []*CommonConfigItem
	resp := svcCtx.DB.Model("ga_common_config_item").Where(whereMap).OrderBy("sort_order asc, id desc").Paginate(ctx, int(req.Page), int(req.PageSize), &list)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return ga.Map{
		"items":     list,
		"total":     resp.Total,
		"page":      resp.Page,
		"page_size": resp.Size,
	}, nil
}

func GetConfigByCategory(ctx context.Context, svcCtx *svc.ServiceContext, categoryKey string) ([]*CommonConfigItem, error) {
	var list []*CommonConfigItem
	resp := svcCtx.DB.Model("ga_common_config_item").
		Where("category_key = ?", categoryKey).
		Where("status = ?", 1).
		OrderBy("sort_order asc, id desc").
		Select(ctx, &list)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	return list, nil
}

func GetConfigDetail(ctx context.Context, svcCtx *svc.ServiceContext, id uint64) (*CommonConfigItem, error) {
	var item CommonConfigItem
	resp := svcCtx.DB.Model("ga_common_config_item").Where("id = ?", id).Find(ctx, &item)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	if resp.IsEmpty() {
		return nil, nil
	}
	return &item, nil
}

func GetConfigByKey(ctx context.Context, svcCtx *svc.ServiceContext, categoryKey, configKey string) (*CommonConfigItem, error) {
	var item CommonConfigItem
	resp := svcCtx.DB.Model("ga_common_config_item").
		Where("category_key = ?", categoryKey).
		Where("config_key = ?", configKey).
		Find(ctx, &item)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	if resp.IsEmpty() {
		return nil, nil
	}
	return &item, nil
}

func CreateConfig(ctx context.Context, svcCtx *svc.ServiceContext, req *configcenter.CreateConfigRequest) (uint64, error) {
	data := map[string]interface{}{
		"category_key":    req.CategoryKey,
		"config_key":      req.ConfigKey,
		"config_name":     req.ConfigName,
		"config_type":     req.ConfigType,
		"config_value":    req.ConfigValue,
		"default_value":   req.DefaultValue,
		"description":     req.Description,
		"options":         req.Options,
		"validation_rule": req.ValidationRule,
		"placeholder":     req.Placeholder,
		"is_required":     req.IsRequired,
		"is_secret":       req.IsSecret,
		"sort_order":      req.SortOrder,
		"status":          req.Status,
	}
	resp := svcCtx.DB.Model("ga_common_config_item").Data(data).Insert(ctx)
	if resp.GetError() != nil {
		return 0, resp.GetError()
	}
	return ga.Uint64(resp.GetLastId()), nil
}

func UpdateConfig(ctx context.Context, svcCtx *svc.ServiceContext, req *configcenter.UpdateConfigRequest) error {
	data := map[string]interface{}{
		"category_key":    req.CategoryKey,
		"config_key":      req.ConfigKey,
		"config_name":     req.ConfigName,
		"config_type":     req.ConfigType,
		"config_value":    req.ConfigValue,
		"default_value":   req.DefaultValue,
		"description":     req.Description,
		"options":         req.Options,
		"validation_rule": req.ValidationRule,
		"placeholder":     req.Placeholder,
		"is_required":     req.IsRequired,
		"is_secret":       req.IsSecret,
		"sort_order":      req.SortOrder,
		"status":          req.Status,
	}
	resp := svcCtx.DB.Model("ga_common_config_item").Where("id = ?", req.Id).Data(data).Update(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

func DeleteConfig(ctx context.Context, svcCtx *svc.ServiceContext, id uint64) error {
	resp := svcCtx.DB.Model("ga_common_config_item").Where("id = ?", id).Delete(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

func UpdateConfigStatus(ctx context.Context, svcCtx *svc.ServiceContext, id uint64, status int32) error {
	resp := svcCtx.DB.Model("ga_common_config_item").Where("id = ?", id).Data(map[string]interface{}{"status": status}).Update(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

func UpdateConfigValue(ctx context.Context, svcCtx *svc.ServiceContext, id uint64, configValue string) error {
	resp := svcCtx.DB.Model("ga_common_config_item").Where("id = ?", id).Data(map[string]interface{}{"config_value": configValue}).Update(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

func SaveConfigValue(ctx context.Context, svcCtx *svc.ServiceContext, categoryKey, configKey, configValue string) error {
	resp := svcCtx.DB.Model("ga_common_config_item").
		Where("category_key = ?", categoryKey).
		Where("config_key = ?", configKey).
		Data(map[string]interface{}{"config_value": configValue}).
		Update(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

func GetConfigValue(ctx context.Context, svcCtx *svc.ServiceContext, categoryKey string) (map[string]string, error) {
	var items []*CommonConfigItem
	resp := svcCtx.DB.Model("ga_common_config_item").
		Fields("config_key,config_value").
		Where("category_key = ?", categoryKey).
		Where("status = ?", 1).
		Select(ctx, &items)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}

	result := make(map[string]string)
	for _, item := range items {
		if item.ConfigValue.Valid {
			result[item.ConfigKey] = item.ConfigValue.String
		} else {
			result[item.ConfigKey] = ""
		}
	}
	return result, nil
}
