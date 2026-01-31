package user

import (
	validate "admin/internal/validate/user"
	"github.com/dwrui/go-zero-admin/pkg/utils/ga"
	"google.golang.org/grpc/status"
	"net/http"

	"admin/internal/logic/user"
	"admin/internal/svc"
	"admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		//解析参数
		if err := ga.ResData(r, &req); err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(err.Error()))
			return
		}
		//验证参数
		if msg := validate.LoginValidate(req); msg != "" {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusOK, ga.Failed().SetMsg(msg))
			return
		}
		// 获取客户端IP和用户代理
		reqs := ga.Map{"ip": ga.GetIp(r), "user_agent": r.UserAgent()}
		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req, reqs)
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
