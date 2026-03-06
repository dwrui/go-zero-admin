package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/model"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpCodeTableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpCodeTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpCodeTableLogic {
	return &UpCodeTableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpCodeTableLogic) UpCodeTable(in *develop.UpCodeTableRequest) (*develop.UpCodeTableResponse, error) {
	tablenames := in.Tablenames

	err := model.UpCodeTable(l.ctx, l.svcCtx, tablenames)
	if err != nil {
		l.Error("更新数据表失败: ", err)
		return &develop.UpCodeTableResponse{}, err
	}

	return &develop.UpCodeTableResponse{}, nil
}
