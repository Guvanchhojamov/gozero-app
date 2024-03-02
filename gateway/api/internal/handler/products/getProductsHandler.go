package products

import (
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/handler/response"
	"net/http"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/logic/products"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := products.NewGetProductsLogic(r.Context(), svcCtx)
		resp, err := l.GetProducts(&req)
		if err != nil {
			response.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("errGetProduct: %s", err.Error()), w)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
