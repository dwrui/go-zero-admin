// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"net/http"

	"admin/internal/logic/config"
	"admin/internal/svc"
	"admin/internal/types"
	validate "admin/internal/validate/config"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func GetConfigValueHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetConfigValueReq
		//解析参数
		if err := ga.ResData(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			return
		}
		//验证参数
		if msg := validate.GetConfigValueValidate(req); msg != "" {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(msg))
			return
		}
		l := config.NewGetConfigValueLogic(r.Context(), svcCtx)
		resp, err := l.GetConfigValue(&req)
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
