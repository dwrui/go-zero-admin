package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/model"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDbFieldLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDbFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDbFieldLogic {
	return &GetDbFieldLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDbFieldLogic) GetDbField(in *develop.GetDbfieldRequest) (*develop.GetDbfieldResponse, error) {
	tablename := in.Tablename
	if tablename == "" {
		return &develop.GetDbfieldResponse{}, nil
	}

	fieldDataList, err := model.GetDbField(l.ctx, l.svcCtx, tablename)
	if err != nil {
		l.Error("获取数据库字段失败: ", err)
		return &develop.GetDbfieldResponse{}, err
	}

	var list []*develop.GetDbFieldData
	for _, field := range fieldDataList {
		list = append(list, &develop.GetDbFieldData{
			Label: field["label"].(string),
			Type:  field["type"].(string),
			Value: field["value"].(string),
		})
	}

	return &develop.GetDbfieldResponse{
		List: list,
	}, nil
}
