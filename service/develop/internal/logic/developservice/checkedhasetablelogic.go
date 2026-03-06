package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/model"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckedHaseTableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckedHaseTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckedHaseTableLogic {
	return &CheckedHaseTableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckedHaseTableLogic) CheckedHaseTable(in *develop.CheckedHaseTableRequest) (*develop.CheckedHaseTableResponse, error) {
	tablenames := in.Tablenames
	_, err := model.CheckedHaseTable(l.ctx, l.svcCtx, tablenames)
	if err != nil {
		l.Error("检查表失败: ", err)
		return &develop.CheckedHaseTableResponse{}, err
	}
	return &develop.CheckedHaseTableResponse{}, nil
}
