package authservicelogic

import (
	"auth/auth"
	"auth/internal/model"
	"auth/internal/svc"
	"context"
	"errors"
	"fmt"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/jwt"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckTokenLogic) CheckToken(in *auth.CheckTokenRequest) (*auth.CheckTokenResponse, error) {
	token := in.Token
	// 生成JWT token
	jwtConfig := jwt.JwtConfig{
		AccessSecret: l.svcCtx.Config.Jwt.AccessSecret,
		AccessExpire: l.svcCtx.Config.Jwt.AccessExpire,
	}
	claimInfo, err := jwt.ParseToken(jwtConfig, token)
	if err != nil {
		return nil, err
	}
	userId := claimInfo.Data["id"]
	exp := claimInfo.Data["exp"]

	//检测是否在token黑名单
	if isBlackList, _ := IsTokenBlackList(l.ctx, l.svcCtx, in.Token); isBlackList {
		return nil, errors.New("token已过期，请重新登录")
	}
	//检验是否是自己系统的用户
	userInfo, err := model.GetUserInfo(l.ctx, l.svcCtx, ga.Uint64(userId), "business_id")
	if err != nil {
		return nil, err
	}
	permission, err := l.svcCtx.Redis.GetCtx(l.ctx, "user_permission:"+ga.String(userId))
	if err != nil {
		return nil, errors.New("数据解析错误")
	}
	var permissionData map[string]interface{}
	// 将字符串解析为 JSON
	if permission != "" {

		if err := json.Unmarshal([]byte(permission), &permissionData); err != nil {
			return nil, errors.New("数据解析错误")
		}
	} else {
		//如果缓存获取不到数据数据库获取并缓存数据
		permissionData, err = model.SetUserPermission(l.ctx, l.svcCtx, ga.Uint64(userId), ga.Uint64(userInfo.BusinessId))
		if err != nil {
			return nil, err
		}
	}
	//判断权限
	permissionDList, ok := permissionData["permissions"].([]interface{})
	if !ok {
		return nil, errors.New("权限数据解析错误")
	}
	if !ga.IsContain(permissionDList, in.Permission) {
		roles, ok := permissionData["roles"].(string)
		if !ok {
			return nil, errors.New("角色数据解析错误")
		}
		if roles != "*" {
			return nil, errors.New("权限不足")
		}
	}

	//判断是否需要刷新token
	expTime, err := time.Parse("2006-01-02 15:04:05", ga.String(exp))
	if err != nil {
		return nil, err
	}
	if expTime.Unix()-time.Now().Unix() <= l.svcCtx.Config.Jwt.RefreshThreshold {
		//需要刷新token
		//加入黑名单
		AddTokenBlackList(l.ctx, l.svcCtx, token, ga.Int(l.svcCtx.Config.Jwt.AccessExpire))
		//刷新token
		newToken, err := jwt.GenerateToken(jwtConfig, ga.Map{"id": claimInfo.Data["id"], "account_id": claimInfo.Data["account_id"], "business_id": claimInfo.Data["business_id"]})
		if err != nil {
			return nil, err
		}
		return &auth.CheckTokenResponse{
			NewToken: newToken,
		}, nil
	}
	return &auth.CheckTokenResponse{}, nil
}

func IsTokenBlackList(ctx context.Context, svcCtx *svc.ServiceContext, token string) (bool, error) {
	key := fmt.Sprintf("token:blacklist:%s", token)

	exists, err := svcCtx.Redis.ExistsCtx(ctx, key)
	if err != nil {
		return false, fmt.Errorf("检查黑名单失败: %v", err)
	}

	return exists, nil
}

func AddTokenBlackList(ctx context.Context, svcCtx *svc.ServiceContext, token string, ttl int) error {
	key := fmt.Sprintf("token:blacklist:%s", token)

	if err := svcCtx.Redis.SetexCtx(ctx, key, "1", ttl); err != nil {
		return fmt.Errorf("保存黑名单条目失败: %v", err)
	}
	return nil
}
