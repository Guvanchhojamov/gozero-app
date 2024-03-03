package orders

import (
	"net/http"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/logic/orders"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOrderByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := orders.NewGetOrderByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
