package common

import (
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"net/http"

	"admin/internal/logic/common"
	"admin/internal/svc"
	"admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := common.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Success().SetData(resp))
		}
	}
}
