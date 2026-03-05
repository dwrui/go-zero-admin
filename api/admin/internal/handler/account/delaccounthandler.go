package account

import (
	"net/http"

	"admin/internal/logic/account"
	"admin/internal/svc"
	"admin/internal/types"
	validate "admin/internal/validate/system"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func DelAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelAccountReq
		if err := ga.ResData(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			return
		}
		if msg := validate.DelAccountValidate(req); msg != "" {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(msg))
			return
		}

		l := account.NewDelAccountLogic(r.Context(), svcCtx)
		resp, err := l.DelAccount(&req)
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