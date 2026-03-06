package developservicelogic

import (
	"context"
	"develop/internal/model"

	"develop/develop"
	"develop/internal/svc"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gconv"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gmap"
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

func (l *GetListLogic) GetList(in *develop.GetListRequest) (*develop.GetDevelopListResponse, error) {
	if in.Page == 0 {
		in.Page = 1
	}
	if in.PageSize == 0 {
		in.PageSize = 10
	}
	whereMap := gmap.New()
	if in.Name != "" {
		whereMap.Set("comment like ? OR cate_tablename  like ?", ga.Slice{"%" + gconv.String(in.Name) + "%", "%" + gconv.String(in.Name) + "%"})
	}
	if in.Status != "" {
		whereMap.Set("status = ?", in.Status)
	}
	if in.CreateTime != "" {
		datetime_arr := ga.SplitAndStr(in.CreateTime, ",")
		whereMap.Set("c.create_time between ? and ?", ga.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	// 调用模型层获取数据
	data, err := model.GetDevelopList(l.ctx, l.svcCtx, whereMap, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	var DevelopListResponse []*develop.DevelopListData
	list := data["list"].([]*model.CommonGenerateCodeModel)
	for _, codeModel := range list {
		DevelopListResponse = append(DevelopListResponse, &develop.DevelopListData{
			ApiFilename:   codeModel.ApiFilename,
			ApiPath:       codeModel.ApiPath,
			AutoIncrement: ga.Uint64(codeModel.AutoIncrement),
			CateTablename: codeModel.CateTablename,
			Codelocation:  codeModel.Codelocation,
			Comment:       codeModel.Comment,
			Component:     codeModel.Component,
			Engine:        codeModel.Engine,
			Fields:        codeModel.Fields,
			Fromtype:      ga.Uint64(codeModel.Fromtype),
			Godir:         codeModel.Godir,
			Icon:          codeModel.Icon,
			Id:            codeModel.Id,
			IsInstall:     ga.Uint64(codeModel.IsInstall),
			Pid:           ga.Uint64(codeModel.Pid),
			Routename:     codeModel.Routename,
			Routepath:     codeModel.Routepath,
			RuleId:        ga.Uint64(codeModel.RuleId),
			RuleName:      codeModel.RuleName,
			Status:        ga.Uint64(codeModel.Status),
			TableRows:     ga.Uint64(codeModel.TableRows),
			Tablename:     codeModel.Tablename,
			TplType:       codeModel.TplType,
			Collation:     codeModel.Collation,
			CreateTime:    codeModel.Createtime.Time.Format("2006-01-02 15:04:05"),
			UpdateTime:    codeModel.Updatetime.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return &develop.GetDevelopListResponse{
		Data:     DevelopListResponse,
		Total:    ga.Uint64(data["total"]),
		Page:     ga.Uint64(data["page"]),
		PageSize: ga.Uint64(data["page_size"]),
	}, nil
}
