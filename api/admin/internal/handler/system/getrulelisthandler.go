// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"net/http"

	"admin/internal/logic/system"
	"admin/internal/svc"
	"admin/internal/types"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func GetRuleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRuleListReq
		//解析参数
		if err := ga.ResData(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			return
		}
		l := system.NewGetRuleListLogic(r.Context(), svcCtx)
		resp, err := l.GetRuleList(&req)
		if err != nil {
			if st, ok := status.FromError(err); ok {
				httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(st.Message()))
			} else {
				httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			}
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Success().SetData(resp.Data))
		}
	}
}
