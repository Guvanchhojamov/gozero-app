package authorization

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/handler/response"
	"net/http"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/logic/authorization"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUpReq
		if err := httpx.Parse(r, &req); err != nil {
			response.NewErrorResponse(http.StatusBadRequest, "errInvalidArgumentSignUpHandler", w)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authorization.NewSignUpLogic(r.Context(), svcCtx)
		resp, err := l.SignUp(&req)
		if err != nil {
			response.NewErrorResponse(http.StatusInternalServerError, "errAuthSignupHandler", w)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
