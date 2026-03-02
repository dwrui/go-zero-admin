// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package dept

import (
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpStatusDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpStatusDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpStatusDeptLogic {
	return &UpStatusDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpStatusDeptLogic) UpStatusDept(req *types.UpStatusDeptReq) (resp *types.UpStatusDeptResp, err error) {
	// todo: add your logic here and delete this line

	return
}
