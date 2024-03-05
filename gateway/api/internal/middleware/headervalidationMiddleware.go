package middleware

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/app"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/handler/response"
	"github.com/go-errors/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"strings"
)

type HeaderValidationMiddleware struct {
	auth app.Authorization
}

func NewHeaderValidationMiddleware(auth app.Authorization) *HeaderValidationMiddleware {
	return &HeaderValidationMiddleware{auth: auth}
}

func (m *HeaderValidationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get Bearer token
		authHeader := r.Header.Get(authHeaderKey)
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			response.NewErrorResponse(http.StatusBadRequest, "invalid bearer token", w)
			logx.Errorf("invalid bearer token: %s", bearerToken)
			return
		}
		// end
		userId, err := m.auth.ParseToken(bearerToken[1])
		if err != nil {
			response.NewErrorResponse(http.StatusUnauthorized, errors.New("invalid token").Error(), w)
			logx.Errorf("invalid token: %s", bearerToken[1])
			return
		}
		ctx := r.Context()
		ctx = metadata.AppendToOutgoingContext(r.Context(), userIdCtxKey, strconv.FormatUint(uint64(userId), 10))
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
