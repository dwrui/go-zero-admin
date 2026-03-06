package model

import (
	"context"
	"database/sql"
	"develop/internal/svc"
	"strings"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gvar"
)

type CommonGenerateCodeField struct {
	Id             uint64       `db:"id"`              // ID
	GeneratecodeId int64        `db:"generatecode_id"` // 关联列表
	Islist         int64        `db:"islist"`          // 是否是列表1=是
	Name           string       `db:"name"`            // 字段名称
	Field          string       `db:"field"`           // 字段
	Isorder        int64        `db:"isorder"`         // 是否参与排序
	Align          string       `db:"align"`           // 对齐方向
	Width          int64        `db:"width"`           // 宽度
	ShowUi         string       `db:"show_ui"`         // 显示UI
	Isform         int64        `db:"isform"`          // 是否为表单字段
	Required       int64        `db:"required"`        // 是否为必填项
	Formtype       string       `db:"formtype"`        // 表单类型
	Datatable      string       `db:"datatable"`       // 关联数据表
	Datatablename  string       `db:"datatablename"`   // 关联显示字段
	DicGroupId     int64        `db:"dic_group_id"`    // 关联字典分组id
	Issearch       int64        `db:"issearch"`        // 是否查询
	Searchway      string       `db:"searchway"`       // 查询方式
	Searchtype     string       `db:"searchtype"`      // 查询文本类型
	FieldWeigh     int64        `db:"field_weigh"`     // 表单排序
	ListWeigh      int64        `db:"list_weigh"`      // 列表排序
	SearchWeigh    int64        `db:"search_weigh"`    // 搜索排序
	DefValue       string       `db:"def_value"`       // 默认值
	OptionValue    string       `db:"option_value"`    // 选项值
	Gridwidth      int64        `db:"gridwidth"`       // 布局栅格
	Searchwidth    int64        `db:"searchwidth"`     // 搜索表单宽
	Createtime     sql.NullTime `db:"createtime"`      // 创建时间
	Updatetime     sql.NullTime `db:"updatetime"`      // 修改时间
}

// TableFieldInfo 数据库表字段信息结构体
type TableFieldInfo struct {
	COLUMN_NAME              string `db:"COLUMN_NAME"`              // 字段名
	COLUMN_COMMENT           string `db:"COLUMN_COMMENT"`           // 字段注释
	DATA_TYPE                string `db:"DATA_TYPE"`                // 数据类型
	CHARACTER_MAXIMUM_LENGTH int64  `db:"CHARACTER_MAXIMUM_LENGTH"` // 字符最大长度
	COLUMN_DEFAULT           string `db:"COLUMN_DEFAULT"`           // 默认值
}

func GetTableFieldInfo(ctx context.Context, svcCtx *svc.ServiceContext, tablename string, id uint64) ([]map[string]interface{}, error) {
	// 获取数据库名
	dbName := svcCtx.Config.Mysql.Database

	// 设置 SQL 模式
	_, err := svcCtx.DB.Exec(ctx, "SET @@sql_mode='NO_ENGINE_SUBSTITUTION';")
	if err != nil {
		return nil, err
	}
	var datalist []*TableFieldInfo
	var dielddata_list []ga.Map
	var haseids []interface{}
	err = svcCtx.DB.Query(ctx, &datalist, "SELECT COLUMN_NAME, COLUMN_COMMENT, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH, COLUMN_DEFAULT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? ORDER BY ORDINAL_POSITION", dbName, tablename)
	if err != nil {
		return nil, err
	}
	for _, data := range datalist {
		if data.COLUMN_COMMENT == "" && data.COLUMN_NAME == "id" {
			data.COLUMN_COMMENT = "ID"
		}
		formtype := "text"
		gridwidth := 12
		width := 100
		show_ui := ""
		isorder := 0
		issearch := 0
		searchway := "="
		searchtype := "text"
		if data.COLUMN_NAME == "id" {
			isorder = 1
			width = 80
		}
		if strings.HasSuffix(data.COLUMN_NAME, "date") {
			formtype = "date"
			width = 160
		} else if strings.HasSuffix(data.COLUMN_NAME, "datetime") {
			formtype = "datetime"
			width = 160
		} else if strings.HasSuffix(data.COLUMN_NAME, "time") {
			formtype = "time"
			width = 120
		} else if strings.HasSuffix(data.COLUMN_NAME, "color") {
			formtype = "colorpicker"
			show_ui = "color"
			width = 80
		} else if strings.HasSuffix(data.COLUMN_NAME, "image") {
			gridwidth = 24
			formtype = "image"
			show_ui = "image"
			width = 80
		} else if strings.HasSuffix(data.COLUMN_NAME, "images") {
			formtype = "images"
			show_ui = "images"
			gridwidth = 24
		} else if strings.HasSuffix(data.COLUMN_NAME, "audio") {
			formtype = "audio"
			gridwidth = 24
		} else if strings.HasSuffix(data.COLUMN_NAME, "file") {
			formtype = "file"
			gridwidth = 24
		} else if strings.HasSuffix(data.COLUMN_NAME, "files") {
			formtype = "files"
			gridwidth = 24
		} else if (strings.HasSuffix(data.COLUMN_NAME, "gender") || strings.HasSuffix(data.COLUMN_NAME, "sex")) && data.DATA_TYPE == "tinyint" {
			formtype = "radio"
			show_ui = "gender"
			width = 110
		} else if data.DATA_TYPE == "int" {
			formtype = "number"
		} else if data.DATA_TYPE == "varchar" && data.CHARACTER_MAXIMUM_LENGTH <= 50 {
			width = 190
		} else if data.DATA_TYPE == "varchar" && data.CHARACTER_MAXIMUM_LENGTH > 50 && data.CHARACTER_MAXIMUM_LENGTH < 225 {
			width = 250
		} else if data.DATA_TYPE == "varchar" && data.CHARACTER_MAXIMUM_LENGTH >= 225 {
			formtype = "textarea"
			width = 280
			show_ui = "des"
			gridwidth = 24
		} else if data.DATA_TYPE == "text" || data.DATA_TYPE == "longtext" {
			formtype = "editor"
			show_ui = "des"
		} else if data.DATA_TYPE == "enum" {
			formtype = "select"
			show_ui = "tags"
			searchtype = "select"
		} else if data.DATA_TYPE == "tinyint" {
			formtype = "radio"
			show_ui = "tag"
		}
		//备注
		option_value := ""
		name_value := data.COLUMN_COMMENT
		if strings.Contains(name_value, ":") {
			name_arr := strings.Split(name_value, ":")
			name_value = name_arr[0]
			option_value = name_arr[1]
		}
		fieldval := svcCtx.DB.Model("common_generatecode_field").Where("generatecode_id", id).Where("field", data.COLUMN_NAME).Value(ctx, "id")
		if fieldval.IsNotEmpty() {
			haseids = append(haseids, fieldval.GetData())
			if option_value != "" {
				dielddata_list = append(dielddata_list, ga.Map{"id": fieldval, "option_value": option_value})
			}
		} else {
			if data.COLUMN_DEFAULT == "" {
				data.COLUMN_DEFAULT = ""
			}
			if data.DATA_TYPE == "varchar" && (data.COLUMN_NAME == "name" || data.COLUMN_NAME == "title") {
				issearch = 1
			}
			if data.DATA_TYPE == "varchar" && data.COLUMN_NAME == "status" {
				issearch = 1
			}
			if data.DATA_TYPE == "create_time" {
				issearch = 1
			}
			if data.DATA_TYPE == "create_time" || data.COLUMN_NAME == "update_time" || data.COLUMN_NAME == "delete_time" {
				searchway = "between"
				searchtype = "daterange"
			}
			maxid := svcCtx.DB.Model("common_generatecode_field").Where("generatecode_id", id).OrderByDesc("id").Value(ctx, "id")
			dielddata_list = append(dielddata_list, ga.Map{"generatecode_id": id, "name": name_value, "option_value": option_value, "field": data.COLUMN_NAME, "formtype": formtype, "gridwidth": gridwidth, "def_value": data.COLUMN_DEFAULT, "isorder": isorder, "issearch": issearch, "searchway": searchway, "searchtype": searchtype, "field_weigh": maxid.GetData(), "list_weigh": maxid.GetData(), "search_weigh": maxid.GetData(), "width": width, "show_ui": show_ui})
		}
	}
	if haseids != nil {
		svcCtx.DB.Model("common_generatecode_field").Where("generatecode_id", id).WhereNotIn("id", haseids).Delete(ctx)
	}
	if dielddata_list != nil {
		resp := svcCtx.DB.Model("common_generatecode_field").Data(dielddata_list).Save(ctx)
		if resp.GetError() != nil {
			return nil, resp.GetError()
		}
		return dielddata_list, nil
	}
	return nil, nil
}

func GetFieldAllList(ctx context.Context, svcCtx *svc.ServiceContext, generatecode_id uint64) (ga.Map, error) {
	var fieldList []*CommonGenerateCodeField
	fieldListResp := svcCtx.DB.Model("common_generatecode_field").Where("generatecode_id", generatecode_id).Order("field_weigh", "asc").Order("id", "asc").Select(ctx, &fieldList)
	if fieldListResp.GetError() != nil {
		return nil, fieldListResp.GetError()
	}
	var fieldListMap []ga.Map
	var listMap []ga.Map
	var SearchListMap []ga.Map
	for _, field := range fieldList {
		fieldData := ga.Map{
			"id":            field.Id,
			"name":          field.Name,
			"field":         field.Field,
			"formtype":      field.Formtype,
			"datatable":     field.Datatable,
			"datatablename": field.Datatablename,
			"dic_group_id":  field.DicGroupId,
			"field_weigh":   field.FieldWeigh,
			"gridwidth":     field.Gridwidth,
		}
		if field.Isform == 1 {
			fieldData["isform"] = gvar.New(true)
		} else {
			fieldData["isform"] = gvar.New(false)
		}
		if field.Required == 1 {
			fieldData["required"] = gvar.New(true)
		} else {
			fieldData["required"] = gvar.New(false)
		}
		if strings.Contains(field.Name, ":") {
			fieldData["name"] = gvar.New(strings.Split(field.Name, ":")[0])
		}
		fieldListMap = append(fieldListMap, fieldData)
		listData := ga.Map{
			"id":         field.Id,
			"name":       field.Name,
			"field":      field.Field,
			"isorder":    field.Isorder,
			"align":      field.Align,
			"width":      field.Width,
			"show_ui":    field.ShowUi,
			"list_weigh": field.ListWeigh,
		}
		if field.Islist == 1 {
			listData["islist"] = gvar.New(true)
		} else {
			listData["islist"] = gvar.New(false)
		}
		if field.Isorder == 1 {
			listData["isorder"] = gvar.New(true)
		} else {
			listData["isorder"] = gvar.New(false)
		}
		listMap = append(listMap, listData)
		saerchData := ga.Map{
			"id":           field.Id,
			"name":         field.Name,
			"searchway":    field.Searchway,
			"searchtype":   field.Searchtype,
			"search_weigh": field.SearchWeigh,
			"searchwidth":  field.Searchwidth,
		}
		if field.Issearch == 1 {
			saerchData["issearch"] = gvar.New(true)
		} else {
			saerchData["issearch"] = gvar.New(false)
		}
		SearchListMap = append(SearchListMap, saerchData)
	}
	return ga.Map{"fieldListMap": fieldListMap, "listMap": listMap, "SearchListMap": SearchListMap}, nil
}
