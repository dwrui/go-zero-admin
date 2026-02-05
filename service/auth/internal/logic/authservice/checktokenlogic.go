package authservicelogic

import (
	"context"
	"errors"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/dwrui/go-zero-admin/pkg/utils/jwt"
	"github.com/dwrui/go-zero-admin/pkg/utils/tools/json"

	"auth/auth"
	"auth/internal/svc"

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
	// todo: 校验token是否有效
	jwtConfig := jwt.JwtConfig{
		AccessSecret: l.svcCtx.Config.Jwt.AccessSecret,
		AccessExpire: l.svcCtx.Config.Jwt.AccessExpire,
	}
	claims, err := jwt.ParseToken(jwtConfig, in.Token)
	if err != nil {
		return nil, err
	}
	user_id := claims.Data["id"]
	permission, err := l.svcCtx.Redis.GetCtx(l.ctx, "user_permission:"+ga.String(user_id))
	if err != nil {
		return nil, errors.New("数据解析错误")
	}
	// 将字符串解析为 JSON
	var permissionData map[string]interface{}
	if err := json.Unmarshal([]byte(permission), &permissionData); err != nil {
		return nil, errors.New("数据解析错误")
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
	return &auth.CheckTokenResponse{
		UserId: ga.Uint64(user_id),
	}, nil
}
