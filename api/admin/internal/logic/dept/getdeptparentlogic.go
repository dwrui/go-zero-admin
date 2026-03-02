// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package dept

import (
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeptParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDeptParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeptParentLogic {
	return &GetDeptParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDeptParentLogic) GetDeptParent(req *types.GetDeptParentReq) (resp *types.GetDeptParentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
