// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package rule

import (
	"context"

	"admin/grpc-client/system"
	"admin/internal/svc"
	"admin/internal/types"

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

func (l *GetRuleListLogic) GetRuleList(req *types.GetRuleListReq) (resp *types.GetRuleListResp, err error) {
	rpcResp, err := l.svcCtx.SystemRuleClient.GetList(l.ctx, &system.GetRuleListRequest{})
	if err != nil {
		return nil, err
	}

	data := make([]types.RuleListData, 0, len(rpcResp.Data))
	for _, item := range rpcResp.Data {
		data = append(data, convertRuleListData(item))
	}

	return &types.GetRuleListResp{
		Data: data,
	}, nil
}

func convertRuleListData(item *system.RuleListData) types.RuleListData {
	children := make([]types.RuleListData, 0, len(item.Children))
	for _, child := range item.Children {
		children = append(children, convertRuleListData(child))
	}

	return types.RuleListData{
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
	}
}
