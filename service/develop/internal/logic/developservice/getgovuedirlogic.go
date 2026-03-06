package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoVueDirLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoVueDirLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoVueDirLogic {
	return &GetGoVueDirLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoVueDirLogic) GetGoVueDir(in *develop.GetGoVueDirRequest) (*develop.GetGoVueDirResponse, error) {
	// todo: add your logic here and delete this line

	return &develop.GetGoVueDirResponse{}, nil
}
