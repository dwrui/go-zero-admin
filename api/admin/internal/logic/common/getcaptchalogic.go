package common

import (
	"common/common"
	"context"
	"fmt"

	"admin/internal/svc"
	"admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(req *types.GetCaptchaReq) (any, error) {
	// todo: add your logic here and delete this line

	rpcResp, err := l.svcCtx.GreetClient.GetCaptcha(l.ctx, &common.GetCaptchaRequest{
		// 设置请求参数
		Type: req.Captcha_type,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(rpcResp)
	return rpcResp, nil
}
