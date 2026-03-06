package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UninstallCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUninstallCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UninstallCodeLogic {
	return &UninstallCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UninstallCodeLogic) UninstallCode(in *develop.UninstallCodeRequest) (*develop.UninstallCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &develop.UninstallCodeResponse{}, nil
}
