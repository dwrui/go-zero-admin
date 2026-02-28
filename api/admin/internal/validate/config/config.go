package validate

import (
	"admin/internal/types"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
)

// CreateCategoryValidate 验证创建配置分类请求
func CreateCategoryValidate(req types.CreateCategoryReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"category_key": map[string]string{
			"required": "请输入分类键",
		},
		"category_name": map[string]string{
			"required": "请输入分类名称",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// UpdateCategoryValidate 验证更新配置分类请求
func UpdateCategoryValidate(req types.UpdateCategoryReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "请输入分类ID",
		},
		"category_name": map[string]string{
			"required": "请输入分类名称",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// DeleteCategoryValidate 验证删除配置分类请求
func DeleteCategoryValidate(req types.DeleteCategoryReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "请输入分类ID",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// CreateConfigValidate 验证创建配置项请求
func CreateConfigValidate(req types.CreateConfigReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"category_key": map[string]string{
			"required": "请输入分类键",
		},
		"config_key": map[string]string{
			"required": "请输入配置键",
		},
		"config_name": map[string]string{
			"required": "请输入配置名称",
		},
		"config_type": map[string]string{
			"required": "请输入配置类型",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// UpdateConfigValidate 验证更新配置项请求
func UpdateConfigValidate(req types.UpdateConfigReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "请输入配置项ID",
		},
		"category_key": map[string]string{
			"required": "请输入分类键",
		},
		"config_key": map[string]string{
			"required": "请输入配置键",
		},
		"config_name": map[string]string{
			"required": "请输入配置名称",
		},
		"config_type": map[string]string{
			"required": "请输入配置类型",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// DeleteConfigValidate 验证删除配置项请求
func DeleteConfigValidate(req types.DeleteConfigReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "请输入配置项ID",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// UpdateConfigStatusValidate 验证更新配置项状态请求
func UpdateConfigStatusValidate(req types.UpdateConfigStatusReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "请输入配置项ID",
		},
		"status": map[string]string{
			"required": "请输入状态",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// SaveConfigValueValidate 验证保存配置值请求
func SaveConfigValueValidate(req types.SaveConfigValueReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"category_key": map[string]string{
			"required": "请输入分类键",
		},
		"config_values": map[string]string{
			"required": "请输入配置值",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// GetConfigByCategoryValidate 验证根据分类获取配置项请求
func GetConfigByCategoryValidate(req types.GetConfigByCategoryReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"category_key": map[string]string{
			"required": "请输入分类键",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// GetConfigDetailValidate 验证获取配置项详情请求
func GetConfigDetailValidate(req types.GetConfigDetailReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "请输入配置项ID",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// GetCategoryDetailValidate 验证获取配置分类详情请求
func GetCategoryDetailValidate(req types.GetCategoryDetailReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"id": map[string]string{
			"required": "请输入分类ID",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}

// GetConfigValueValidate 验证获取配置值请求
func GetConfigValueValidate(req types.GetConfigValueReq) string {
	validator := ga.Validator()
	messages := map[string]interface{}{
		"category_key": map[string]string{
			"required": "请输入分类键",
		},
	}
	validator.SetMessages(messages)
	if err := validator.ValidateOne(req); err != nil {
		return err.Error()
	}
	return ""
}