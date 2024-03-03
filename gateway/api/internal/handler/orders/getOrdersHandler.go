package orders

import (
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/handler/response"
	"net/http"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/logic/orders"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrdersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := orders.NewGetOrdersLogic(r.Context(), svcCtx)
		resp, err := l.GetOrders(&req)
		if err != nil {
			response.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("errGetOrders: %s", err.Error()), w)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
