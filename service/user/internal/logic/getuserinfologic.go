package logic

import (
	"context"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserinfoLogic {
	return &GetUserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserinfoLogic) GetUserinfo(in *user.GetUserinfoRequest) (*user.GetUserinfoResponse, error) {
	// todo: add your logic here and delete this line
	userID := in.UserId
	account := model.GaBusinessAccount{}
	userInfo := l.svcCtx.DB.Model("business_account").Where("id = ?", userID).Find(l.ctx, &account)
	if userInfo.GetError() != nil {
		return nil, userInfo.GetError()
	}
	if account.Avatar == "" {
		account.Avatar = "http://localhost:8881/common/static/unknown.png"
	}
	account.Mobile = ga.HideStrInfo("mobile", account.Mobile)
	account.Email = ga.HideStrInfo("email", account.Email)
	return &user.GetUserinfoResponse{
		Id:         account.Id,
		BusinessId: ga.Uint64(account.BusinessId),
		Name:       account.Name,
		Nickname:   account.Nickname,
		Mobile:     account.Mobile,
		Email:      account.Email,
		Avatar:     account.Avatar,
		Status:     ga.Uint64(account.Status),
		Createtime: ga.String(account.CreateTime.Time.Format("2006-01-02 15:04:05")),
	}, nil
}
