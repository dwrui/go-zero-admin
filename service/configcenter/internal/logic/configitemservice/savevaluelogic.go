package configitemservicelogic

import (
	"configcenter/configcenter"
	"configcenter/internal/model"
	"configcenter/internal/svc"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveValueLogic {
	return &SaveValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveValueLogic) SaveValue(in *configcenter.SaveConfigValueRequest) (*configcenter.SaveConfigValueResponse, error) {
	// 检查分类是否存在
	category, err := model.GetCategoryByKey(l.ctx, l.svcCtx, in.CategoryKey)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 遍历保存每个配置值
	for configKey, configValue := range in.ConfigValues {
		// 检查配置项是否存在
		item, err := model.GetConfigByKey(l.ctx, l.svcCtx, in.CategoryKey, configKey)
		if err != nil {
			return nil, err
		}
		if item == nil {
			continue // 跳过不存在的配置项
		}

		// 保存配置值
		err = model.SaveConfigValue(l.ctx, l.svcCtx, in.CategoryKey, configKey, configValue)
		if err != nil {
			return nil, err
		}
	}

	return &configcenter.SaveConfigValueResponse{}, nil
}
