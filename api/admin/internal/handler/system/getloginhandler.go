// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package system

import (
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"google.golang.org/grpc/status"
	"net/http"

	"admin/internal/logic/system"
	"admin/internal/svc"
	"admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetLoginReq
		l := system.NewGetLoginLogic(r.Context(), svcCtx)
		resp, err := l.GetLogin(&req)
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
