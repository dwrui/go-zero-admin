package user

import (
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"google.golang.org/grpc/status"
	"net/http"

	"admin/internal/logic/user"
	"admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		l := user.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(token)
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
