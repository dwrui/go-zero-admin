package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogic) Del(in *develop.DelRequest) (*develop.DelResponse, error) {
	// todo: add your logic here and delete this line

	return &develop.DelResponse{}, nil
}
