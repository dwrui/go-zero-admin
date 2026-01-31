package common

import (
	validate "admin/internal/validate/common"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"google.golang.org/grpc/status"
	"net/http"

	"admin/internal/logic/common"
	"admin/internal/svc"
	"admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCaptchaReq
		//解析参数
		if err := ga.ResData(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			return
		}
		//验证参数
		if msg := validate.GetCaptchaValidate(req); msg != "" {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(msg))
			return
		}
		l := common.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha(&req)
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
