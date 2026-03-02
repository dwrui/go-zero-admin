// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package dept

import (
	"context"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelDeptLogic {
	return &DelDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelDeptLogic) DelDept(req *types.DelDeptReq) (resp *types.DelDeptResp, err error) {
	// todo: add your logic here and delete this line

	return
}
