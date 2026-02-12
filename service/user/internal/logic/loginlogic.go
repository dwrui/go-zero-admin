package logic

import (
	"context"
	"errors"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/jwt"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/gtime"
	"time"
	"user/internal/model"
	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	if in.Username != "" {
		if in.Password == "" {
			return nil, errors.New("密码不能为空")
		}
		var account model.GaBusinessAccount
		res := l.svcCtx.DB.Model("admin_account").Fields("id,account_id,business_id,password,salt,name,status,login_attempts,lock_time").Where("username =? OR email = ?", in.Username, in.Username).Find(l.ctx, &account)
		if res.GetError() != nil {
			return nil, res.GetError()
		}
		if res.IsEmpty() {
			return nil, errors.New("账号不存在!")
		}
		if account.Status != 1 {
			return nil, errors.New("账号已被禁用!")
		}
		if time.Now().Before(account.LockTime) {
			return nil, errors.New("账号已被锁定，请稍后重试!")
		}
		pass := ga.Md5(in.Password + account.Salt)
		//判断密码是否正确
		if pass != account.Password {
			// 添加日志操作
			go model.AddloginLog(l.ctx, l.svcCtx, ga.Map{"uid": account.Id, "account_id": account.AccountId, "business_id": account.BusinessId, "type": "business", "status": 0, "des": "账号登录", "error_msg": "输入的密码不正确！", "ip": in.ClientIp, "user_agent": in.UserAgent})
			if account.LoginAttempts >= 3 {
				//锁定账户
				err := model.LockAccount(l.ctx, l.svcCtx, account.Id)
				if err != nil {
					return nil, errors.New("密码错误次数过多，账户已被锁定30分钟")
				}
			}
			l.svcCtx.DB.Model("admin_account").Where("id = ?", account.Id).Inc(l.ctx, "login_attempts", 1)
			return nil, errors.New("密码错误!")
		}
		//暂时先不做资源共享
		//if !ga.VerifyCaptcha(in.Codeid, in.Captcha) {
		//	model.AddloginLog(l.ctx, l.svcCtx, ga.Map{"uid": account.Id, "account_id": account.AccountId, "business_id": account.BusinessId, "type": "business", "status": 0, "des": "账号登录", "error_msg": "输入的验证码不正确！"})
		//	return nil, errors.New("您输入的验证码不正确!")
		//}
		// 生成JWT token
		jwtConfig := jwt.JwtConfig{
			AccessSecret: l.svcCtx.Config.Jwt.AccessSecret,
			AccessExpire: l.svcCtx.Config.Jwt.AccessExpire,
		}
		token, err := jwt.GenerateToken(jwtConfig, ga.Map{"id": account.Id, "account_id": account.AccountId, "business_id": account.BusinessId})
		if err != nil {
			return nil, err
		}
		go model.AddloginLog(l.ctx, l.svcCtx, ga.Map{"uid": account.Id, "account_id": account.AccountId, "business_id": account.BusinessId, "type": "business", "status": 1, "des": "账号登录", "ip": in.ClientIp, "user_agent": in.UserAgent})
		l.svcCtx.DB.Model("admin_account").Where("id = ?", account.Id).Data(ga.Map{"loginstatus": 1, "last_login_time": gtime.Timestamp(), "last_login_ip": in.ClientIp}).Update(l.ctx)
		return &user.LoginResponse{
			Data: token,
		}, nil
	} else if in.Email != "" {

	} else if in.Mobile != "" {

	} else {

	}
	return &user.LoginResponse{}, nil
}
