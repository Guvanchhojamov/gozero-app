package authorization

import (
	"net/http"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/logic/authorization"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authorization.NewSignInLogic(r.Context(), svcCtx)
		resp, err := l.SignIn(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
