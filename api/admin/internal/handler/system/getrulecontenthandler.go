// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"fmt"
	"net/http"

	"admin/internal/logic/system"
	"admin/internal/svc"
	"admin/internal/types"
	validate "admin/internal/validate/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func GetRuleContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRuleContentReq
		//解析参数
		if err := ga.ResData(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			return
		}
		fmt.Println("GetRuleContentHandler", req)
		//验证参数
		if msg := validate.GetRuleContentValidate(req); msg != "" {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(msg))
			return
		}
		l := system.NewGetRuleContentLogic(r.Context(), svcCtx)
		resp, err := l.GetRuleContent(&req)
		if err != nil {
			if st, ok := status.FromError(err); ok {
				httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(st.Message()))
			} else {
				httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			}
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Success().SetData(resp))
		}
	}
}
