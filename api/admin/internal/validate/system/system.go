package validate

import (
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

// GetLoginValidate 验证获取登录日志请求
func GetLoginValidate(req types.GetLoginReq) string {
	return ""
}

// SaveRuleValidate 验证保存路由请求
func SaveRuleValidate(req types.SaveRuleReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"title": map[string]string{
			"required": "请输入菜单名称",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// UpStatusRuleValidate 验证更新路由状态请求
func UpStatusRuleValidate(req types.UpStatusRuleReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "参数错误",
		},
		"status": map[string]string{
			"required": "参数错误",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// DelRuleValidate 验证删除路由请求
func DelRuleValidate(req types.DelRuleReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "参数或格式错误",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// GetRuleContentValidate 验证获取路由详情请求
func GetRuleContentValidate(req types.GetRuleContentReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "参数错误",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// GetRuleParentValidate 验证获取路由父级请求
func GetRuleParentValidate(req types.GetRuleParentReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "参数错误",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// GetRuleListValidate 验证获取路由列表请求
func GetRuleListValidate(req types.GetRuleListReq) string {
	return ""
}
