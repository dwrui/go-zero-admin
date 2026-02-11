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
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

func SaveQuickValidate(req types.SaveQuickReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"name": map[string]string{
			"required": "请输入快速编辑名称",
		},
		"path_url": map[string]string{
			"required": "请输入快速编辑路径",
		},
		"icon": map[string]string{
			"required": "请输入快速编辑图标",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}
