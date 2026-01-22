package user

import (
	"net/http"

	"admin/internal/logic/user"
	"admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		err := l.GetUserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
