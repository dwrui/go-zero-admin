// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package common

import (
	"net/http"

	"admin/internal/logic/common"
	"admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetQuickHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := common.NewGetQuickLogic(r.Context(), svcCtx)
		err := l.GetQuick()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
