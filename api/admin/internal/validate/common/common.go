package validate

import (
	"admin/internal/types"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

func GetCaptchaValidate(req types.GetCaptchaReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"type": map[string]string{
			"required": "请选择验证码类型",
		},
	}
	validator.SetMessages(messages)
	if err := validator.Validate(req); err != nil {
		return err.Error()
	}
	return ""
}
