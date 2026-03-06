package developservicelogic

import (
	"context"

	"develop/develop"
	"develop/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuParentLogic {
	return &GetMenuParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMenuParentLogic) GetMenuParent(in *develop.GetDevelopMenuParentRequest) (*develop.GetDevelopMenuParentResponse, error) {
	// todo: add your logic here and delete this line

	return &develop.GetDevelopMenuParentResponse{}, nil
}
