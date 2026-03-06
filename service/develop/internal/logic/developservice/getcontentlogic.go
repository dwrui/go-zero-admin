package developservicelogic

import (
	"context"
	"develop/internal/model"
	"os"
	"path/filepath"

	"develop/develop"
	"develop/internal/svc"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gconv"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentLogic) GetContent(in *develop.GetDevelopContentRequest) (*develop.GetDevelopContentResponse, error) {
	// todo: add your logic here and delete this line
	data, err := model.GetContent(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	var dataRequest *develop.DevelopContentData
	dataRequest.Id = data.Id
	dataRequest.Tablename = data.Tablename
	dataRequest.ApiFilename = data.ApiFilename
	dataRequest.ApiPath = data.ApiPath
	dataRequest.CateTablename = data.CateTablename
	dataRequest.Codelocation = data.Codelocation
	dataRequest.Comment = data.Comment
	dataRequest.Component = data.Component
	dataRequest.Icon = data.Icon
	dataRequest.Pid = ga.Uint64(data.Pid)
	dataRequest.IsInstall = ga.Uint64(data.IsInstall)
	dataRequest.Routename = data.Routename
	dataRequest.Routepath = data.Routepath
	dataRequest.RuleId = ga.Uint64(data.RuleId)
	dataRequest.RuleName = data.RuleName
	dataRequest.Tablename = data.Tablename
	dataRequest.TplType = data.TplType
	_, err = model.GetTableFieldInfo(l.ctx, l.svcCtx, data.Tablename, in.Id)
	if err != nil {
		return nil, err
	}
	fieldAllList, err := model.GetFieldAllList(l.ctx, l.svcCtx, in.Id)
	if err != nil {
		return nil, err
	}
	haseadmin := false
	vue_viewsfiles_path_admin := filepath.Join(gconv.String(l.svcCtx.Config.AdminConfig.Vueobjroot), gconv.String(l.svcCtx.Config.AdminConfig.AdminDirName))
	if _, err := os.Stat(vue_viewsfiles_path_admin); err == nil {
		haseadmin = true
	}
	fieldList := []*develop.DevelopFieldList{}
	for _, fieldListMapOne := range fieldAllList["fieldListMap"].([]map[string]interface{}) {
		fieldList = append(fieldList, &develop.DevelopFieldList{
			Id:            ga.Uint64(fieldListMapOne["id"]),
			Isform:        ga.Bool(fieldListMapOne["isform"]),
			Field:         ga.String(fieldListMapOne["field"]),
			Datatable:     ga.String(fieldListMapOne["datatable"]),
			Datatablename: ga.String(fieldListMapOne["datatablename"]),
			DicGroupId:    ga.Uint64(fieldListMapOne["dic_group_id"]),
			Name:          ga.String(fieldListMapOne["name"]),
			FieldWeigh:    ga.Uint64(fieldListMapOne["field_weigh"]),
			Formtype:      ga.String(fieldListMapOne["formtype"]),
			Gridwidth:     ga.Uint64(fieldListMapOne["gridwidth"]),
			Required:      ga.Bool(fieldListMapOne["required"]),
		})
	}
	return &develop.GetDevelopContentResponse{
		Data:      dataRequest,
		Haseadmin: haseadmin,
		FieldList: fieldList,
	}, nil
}
