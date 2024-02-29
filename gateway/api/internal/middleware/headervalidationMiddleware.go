package middleware

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/domain"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/handler/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"strings"
)

const (
	authHeaderKey = "Authorization"
)

type HeaderValidationMiddleware struct {
	auth domain.Authorization
}

func NewHeaderValidationMiddleware(auth domain.Authorization) *HeaderValidationMiddleware {
	return &HeaderValidationMiddleware{auth: auth}
}

func (m *HeaderValidationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.TracerFromContext(r.Context()).Start(r.Context(), "SignUpAuth.Middleware")
		defer span.End()
		// Get Bearer token
		authHeader := r.Header.Get(authHeaderKey)
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			logx.Errorf("invalid bearer: %s", bearerToken)
			response.NewErrorResponse(http.StatusBadRequest, "invalid bearer", w)
			return
		}
		// end

		userId, err := m.auth.ParseToken(bearerToken[1])
		if err != nil {
			logx.Errorf("invalid token: %s", bearerToken[1])
			response.NewErrorResponse(http.StatusUnauthorized, "invalid token: auhtMiddlewarePareseToken", w)
			return
		}
		ctx = metadata.AppendToOutgoingContext(r.Context(), "userId", strconv.FormatUint(uint64(userId), 10))
		newReq := r.WithContext(ctx)
		// Passthrough to next handler if need
		next(w, newReq)
	}
}
