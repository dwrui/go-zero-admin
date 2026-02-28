// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"net/http"

	"admin/internal/logic/config"
	"admin/internal/svc"
	"admin/internal/types"

	"google.golang.org/grpc/status"

	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCategoryListReq
		//解析参数
		if err := ga.ResData(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			return
		}
		l := config.NewGetCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryList(&req)
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
