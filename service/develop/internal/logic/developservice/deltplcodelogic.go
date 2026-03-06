package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTplCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTplCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTplCodeLogic {
	return &DelTplCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTplCodeLogic) DelTplCode(in *develop.DelTplCodeRequest) (*develop.DelTplCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &develop.DelTplCodeResponse{}, nil
}
