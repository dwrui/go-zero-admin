package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkTplCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkTplCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkTplCodeLogic {
	return &MarkTplCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkTplCodeLogic) MarkTplCode(in *develop.MarkTplCodeRequest) (*develop.MarkTplCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &develop.MarkTplCodeResponse{}, nil
}
