// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package common

import (
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"google.golang.org/grpc/status"
	"net/http"

	"admin/internal/logic/common"
	"admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetQuickHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//token := r.Header.Get("Authorization")
		//err := svcCtx.CheckPermission(r.Context(), r, token, "quick:view")
		//if err != nil {
		//	st, ok := status.FromError(err)
		//	if !ok {
		//		httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(st.Message()))
		//		return
		//	}
		//	httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(st.Message()))
		//	return
		//}

		l := common.NewGetQuickLogic(r.Context(), svcCtx)
		resp, err := l.GetQuick()
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
