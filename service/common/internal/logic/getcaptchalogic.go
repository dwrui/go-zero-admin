package logic

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"

	"common/common"
	"common/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(in *common.GetCaptchaRequest) (*common.GetCaptchaResponse, error) {
	// todo: add your logic here and delete this line
	captchaResult, err := ga.GenerateCaptcha(l.svcCtx.App.LoginCaptcha)
	if err != nil {
		return nil, err
	}
	return &common.GetCaptchaResponse{
		Id:         captchaResult.(ga.CaptchaResult).Id,
		Show:       captchaResult.(ga.CaptchaResult).Show,
		Img:        captchaResult.(ga.CaptchaResult).Base64Blog,
		ExpireTime: captchaResult.(ga.CaptchaResult).ExpireTime,
	}, nil
}
