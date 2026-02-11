package logservicelogic

import (
	"context"

	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLastLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLastLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLastLoginLogic {
	return &DelLastLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLastLoginLogic) DelLastLogin(in *system.DelLastLoginRequest) (*system.DelLastLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &system.DelLastLoginResponse{}, nil
}
