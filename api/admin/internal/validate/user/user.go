package validate

import (
	"admin/internal/types"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

func LoginValidate(req types.LoginReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"username": map[string]string{
			"required": "请输入用户名",
		},
		"password": map[string]string{
			"required": "请输入密码",
		},
		"codeid": map[string]string{
			"required": "验证码错误",
		},
		"captcha": map[string]string{
			"required": "验证码错误",
		},
	}
	validator.SetMessages(messages)
	if err := validator.Validate(req); err != nil {
		return err.Error()
	}
	return ""
}
