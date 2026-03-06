package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTableListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTableListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTableListLogic {
	return &GetTableListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTableListLogic) GetTableList(in *develop.GetTableListRequest) (*develop.GetTableListResponse, error) {
	// todo: add your logic here and delete this line

	return &develop.GetTableListResponse{}, nil
}
