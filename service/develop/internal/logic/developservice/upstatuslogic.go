package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/model"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStatusLogic {
	return &UpStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpStatusLogic) UpStatus(in *develop.UpStatusRequest) (*develop.UpStatusResponse, error) {
	id := in.Id
	status := in.Status

	err := model.UpStatus(l.ctx, l.svcCtx, id, status)
	if err != nil {
		l.Error("更新状态失败: ", err)
		return &develop.UpStatusResponse{}, err
	}

	return &develop.UpStatusResponse{}, nil
}
