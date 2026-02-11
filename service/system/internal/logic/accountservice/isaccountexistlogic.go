package accountservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsaccountexistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsaccountexistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsaccountexistLogic {
	return &IsaccountexistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsaccountexistLogic) Isaccountexist(in *system.IsAccountExistRequest) (*system.IsAccountExistResponse, error) {
	// todo: add your logic here and delete this line

	return &system.IsAccountExistResponse{}, nil
}
