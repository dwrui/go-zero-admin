package model

import (
	"context"
	"database/sql"
	"develop/internal/svc"
	"errors"
	"strings"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
)

type CommonGenerateCodeModel struct {
	Id            uint64       `db:"id"`             // ID
	Fromtype      int64        `db:"fromtype"`       // 数据类型0=数据表，1=代码工具
	Tablename     string       `db:"tablename"`      // 表名称
	Comment       string       `db:"comment"`        // 表备注
	Engine        string       `db:"engine"`         // 引擎
	TableRows     int64        `db:"table_rows"`     // 记录数
	Collation     string       `db:"collation"`      // 编码
	AutoIncrement int64        `db:"auto_increment"` // 自增索引
	Status        int64        `db:"status"`         // 状态1=禁用
	Pid           int64        `db:"pid"`            // 菜单上级
	Icon          string       `db:"icon"`           // 图标
	Routepath     string       `db:"routepath"`      // 路由地址
	Routename     string       `db:"routename"`      // 路由名称
	Component     string       `db:"component"`      // 组件路径
	Godir         string       `db:"godir"`          // 后端代码位置
	ApiPath       string       `db:"api_path"`       // 后端业务接口
	ApiFilename   string       `db:"api_filename"`   // 后端文件名
	Fields        string       `db:"fields"`         // 查询字段
	RuleId        int64        `db:"rule_id"`        // 生成菜单id
	RuleName      string       `db:"rule_name"`      // 菜单名称
	Codelocation  string       `db:"codelocation"`   // 生成代码位置
	IsInstall     int64        `db:"is_install"`     // 是否安装0=未安装，1=已安装，2=已卸载
	TplType       string       `db:"tpl_type"`       // 模板类型list=仅一个数据，cate=数据加分类
	CateTablename string       `db:"cate_tablename"` // 分类表名称
	Createtime    sql.NullTime `db:"createtime"`     // 上传时间
	Updatetime    sql.NullTime `db:"updatetime"`     // 修改时间
}

func GetDevelopList(ctx context.Context, svcCtx *svc.ServiceContext, whereMap *gmap.Map, page, size uint64) (ga.Map, error) {
	var list []*CommonGenerateCodeModel
	resp := svcCtx.DB.Model("common_generatecode").Where(whereMap).Order("c.id", "desc").Paginate(ctx, ga.Int(page), ga.Int(size), &list)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return ga.Map{
		"list":  list,
		"total": resp.Total,
		"page":  resp.Page,
		"size":  resp.Size,
	}, nil
}

// GetDbField 获取数据库表字段信息
type DbField struct {
	COLUMN_NAME    string `db:"COLUMN_NAME"`
	COLUMN_COMMENT string `db:"COLUMN_COMMENT"`
	DATA_TYPE      string `db:"DATA_TYPE"`
}

func GetDbField(ctx context.Context, svcCtx *svc.ServiceContext, tablename string) ([]map[string]interface{}, error) {
	tablenameArr := strings.Split(tablename, ",")
	var fieldDataList []map[string]interface{}

	// 获取数据库名
	dbName := svcCtx.Config.Mysql.Database

	for _, val := range tablenameArr {
		var fields []*DbField
		query := "SELECT COLUMN_NAME, COLUMN_COMMENT, DATA_TYPE FROM information_schema.columns WHERE TABLE_SCHEMA=? AND TABLE_NAME=?"
		err := svcCtx.DB.Query(ctx, &fields, query, dbName, val)
		if err != nil {
			return nil, err
		}

		for _, field := range fields {
			label := field.COLUMN_COMMENT
			if label == "" && field.COLUMN_NAME == "id" {
				label = "ID"
			}

			fieldDataList = append(fieldDataList, map[string]interface{}{
				"value": field.COLUMN_NAME,
				"label": label,
				"type":  field.DATA_TYPE,
			})
		}
	}

	return fieldDataList, nil
}

// CheckedHaseTable 检查数据库中是否存在指定的表
func CheckedHaseTable(ctx context.Context, svcCtx *svc.ServiceContext, tablenames []string) ([]string, error) {
	if len(tablenames) == 0 {
		return []string{}, nil
	}

	// 使用 WhereIn 查询存在的表名
	var existingTables []string
	resp := svcCtx.DB.Model("common_generatecode").WhereIn("tablename", tablenames).Column(ctx, "tablename", &existingTables)
	if resp.GetError() != nil {
		return nil, resp.GetError()
	}

	return existingTables, nil
}

// TableInfo 表信息结构体
type TableInfo struct {
	TABLE_NAME      string `db:"TABLE_NAME"`
	TABLE_COMMENT   string `db:"TABLE_COMMENT"`
	ENGINE          string `db:"ENGINE"`
	TABLE_ROWS      int64  `db:"TABLE_ROWS"`
	TABLE_COLLATION string `db:"TABLE_COLLATION"`
	AUTO_INCREMENT  int64  `db:"AUTO_INCREMENT"`
}

// UpCodeTable 更新生成代码的数据表
func UpCodeTable(ctx context.Context, svcCtx *svc.ServiceContext, tablenames []string) error {
	if len(tablenames) == 0 {
		return nil
	}

	// 获取数据库名
	dbName := svcCtx.Config.Mysql.Database

	// 设置 SQL 模式
	_, err := svcCtx.DB.Exec(ctx, "SET @@sql_mode='NO_ENGINE_SUBSTITUTION';")
	if err != nil {
		return err
	}

	// 准备要保存的数据
	var dataList []map[string]interface{}

	for _, val := range tablenames {
		var tableInfos []*TableInfo
		query := "SELECT TABLE_NAME, TABLE_COMMENT, ENGINE, TABLE_ROWS, TABLE_COLLATION, AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA=? AND TABLE_NAME=?"
		err := svcCtx.DB.Query(ctx, &tableInfos, query, dbName, val)
		if err != nil {
			return err
		}

		for _, tableInfo := range tableInfos {
			comment := tableInfo.TABLE_COMMENT
			if comment == "" {
				comment = ""
			}

			data := map[string]interface{}{
				"tablename":      tableInfo.TABLE_NAME,
				"comment":        comment,
				"rule_name":      comment,
				"engine":         tableInfo.ENGINE,
				"table_rows":     tableInfo.TABLE_ROWS,
				"collation":      tableInfo.TABLE_COLLATION,
				"auto_increment": tableInfo.AUTO_INCREMENT,
			}
			dataList = append(dataList, data)
		}
	}

	// 保存数据
	if len(dataList) > 0 {
		resp := svcCtx.DB.Model("common_generatecode").Data(dataList).Save(ctx)
		if resp.GetError() != nil {
			return resp.GetError()
		}
	}

	return nil
}

// UpStatus 更新生成代码的状态
func UpStatus(ctx context.Context, svcCtx *svc.ServiceContext, id uint64, status uint64) error {
	if id == 0 {
		return nil
	}

	// 更新状态
	resp := svcCtx.DB.Model("common_generatecode").Where("id", id).Data(map[string]interface{}{"status": status}).Update(ctx)
	if resp.GetError() != nil {
		return resp.GetError()
	}

	return nil
}

func GetContent(ctx context.Context, svcCtx *svc.ServiceContext, id uint64) (CommonGenerateCodeModel, error) {
	var content CommonGenerateCodeModel
	resp := svcCtx.DB.Model("common_generatecode").Fields("id,tablename,comment,pid,rule_id,rule_name,icon,is_install,routepath,routename,component,api_path,api_filename,cate_tablename,tpl_type,codelocation").Where("id", id).Find(ctx, &content)
	if resp.GetError() != nil {
		return content, resp.GetError()
	}
	if resp.IsEmpty() {
		return content, errors.New("生成数据表不存在")
	}
	return content, nil
}
