package model

import (
	"configcenter/configcenter"
	"configcenter/internal/svc"
	"context"
	"database/sql"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
)

type CommonConfigCategory struct {
	Id           int64          `json:"id" db:"id"`
	CategoryKey  string         `json:"category_key" db:"category_key"`
	CategoryName string         `json:"category_name" db:"category_name"`
	Description  sql.NullString `json:"description" db:"description"`
	SortOrder    int32          `json:"sort_order" db:"sort_order"`
	IsSystem     int32          `json:"is_system" db:"is_system"`
	CreateTime   sql.NullTime   `json:"create_time" db:"create_time"`
	UpdateTime   sql.NullTime   `json:"update_time" db:"update_time"`
}

func GetCategoryList(ctx context.Context, svcCtx *svc.ServiceContext, req *configcenter.GetCategoryListRequest) (ga.Map, error) {
	whereMap := gmap.New()
	if req.CategoryKey != "" {
		whereMap.Set("category_key like ?", "%"+req.CategoryKey+"%")
	}
	if req.CategoryName != "" {
		whereMap.Set("category_name like ?", "%"+req.CategoryName+"%")
	}

	var list []*CommonConfigCategory
	resp := svcCtx.DB.Model("common_config_category").Where(whereMap).OrderBy("sort_order asc, id desc").Paginate(ctx, int(req.Page), int(req.PageSize), &list)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return ga.Map{
		"items":     resp.Items,
		"total":     resp.Total,
		"page":      resp.Page,
		"page_size": resp.Size,
	}, nil
}

func GetCategoryAll(ctx context.Context, svcCtx *svc.ServiceContext) ([]*CommonConfigCategory, error) {
	var list []*CommonConfigCategory
	resp := svcCtx.DB.Model("ga_common_config_category").OrderBy("sort_order asc, id desc").Select(ctx, &list)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	return list, nil
}

func GetCategoryDetail(ctx context.Context, svcCtx *svc.ServiceContext, id uint64) (*CommonConfigCategory, error) {
	var category CommonConfigCategory
	resp := svcCtx.DB.Model("ga_common_config_category").Where("id = ?", id).Find(ctx, &category)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	if resp.IsEmpty() {
		return nil, nil
	}
	return &category, nil
}

func GetCategoryByKey(ctx context.Context, svcCtx *svc.ServiceContext, categoryKey string) (*CommonConfigCategory, error) {
	var category CommonConfigCategory
	resp := svcCtx.DB.Model("ga_common_config_category").Where("category_key = ?", categoryKey).Find(ctx, &category)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}
	if resp.IsEmpty() {
		return nil, nil
	}
	return &category, nil
}

func CreateCategory(ctx context.Context, svcCtx *svc.ServiceContext, req *configcenter.CreateCategoryRequest) (uint64, error) {
	data := map[string]interface{}{
		"category_key":  req.CategoryKey,
		"category_name": req.CategoryName,
		"description":   req.Description,
		"sort_order":    req.SortOrder,
		"is_system":     0,
	}
	resp := svcCtx.DB.Model("ga_common_config_category").Data(data).Insert(ctx)
	if resp.GetError() != nil {
		return 0, resp.GetError()
	}
	return ga.Uint64(resp.GetLastId()), nil
}

func UpdateCategory(ctx context.Context, svcCtx *svc.ServiceContext, req *configcenter.UpdateCategoryRequest) error {
	data := map[string]interface{}{
		"category_name": req.CategoryName,
		"description":   req.Description,
		"sort_order":    req.SortOrder,
	}
	resp := svcCtx.DB.Model("ga_common_config_category").Where("id = ?", req.Id).Data(data).Update(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}

func DeleteCategory(ctx context.Context, svcCtx *svc.ServiceContext, id uint64) error {
	resp := svcCtx.DB.Model("ga_common_config_category").Where("id = ?", id).Delete(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}
	return nil
}
