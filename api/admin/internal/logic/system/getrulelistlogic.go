// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRuleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRuleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRuleListLogic {
	return &GetRuleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRuleListLogic) GetRuleList(req *types.GetRuleListReq) (types.GetRuleListResp, error) {
	resp, err := l.svcCtx.SystemRuleClient.GetList(l.ctx, &system.GetRuleListRequest{})
	if err != nil {
		return types.GetRuleListResp{}, err
	}
	// 递归函数，将RPC的RuleListData转换为API的RuleListData
	var convertRuleListData func([]*system.RuleListData) []types.RuleListData
	convertRuleListData = func(items []*system.RuleListData) []types.RuleListData {
		result := make([]types.RuleListData, 0)
		for _, item := range items {
			children := make([]types.RuleListData, 0)
			if len(item.Children) > 0 {
				children = convertRuleListData(item.Children) // 递归处理children
			}

			result = append(result, types.RuleListData{
				Component:  item.Component,
				CreateTime: item.CreateTime,
				Icon:       item.Icon,
				Id:         item.Id,
				Locale:     item.Locale,
				Path:       item.Path,
				Permission: item.Permission,
				Pid:        item.Pid,
				Spacer:     item.Spacer,
				Status:     item.Status,
				Title:      item.Title,
				Type:       item.Type,
				Weigh:      item.Weigh,
				Children:   children,
			})
		}
		return result
	}
	// 转换数据
	ruleListData := convertRuleListData(resp.Data)
	// 确保即使没有数据时也返回空切片而不是null
	if ruleListData == nil {
		ruleListData = make([]types.RuleListData, 0)
	}
	return types.GetRuleListResp{
		Data: ruleListData,
	}, nil
}
