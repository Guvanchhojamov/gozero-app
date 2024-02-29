package middleware

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/domain"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/app"
	"github.com/go-errors/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
)

type RolePermissionMiddleware struct {
	auth domain.Authorization
}

func NewRolePermissionMiddleware(auth domain.Authorization) *RolePermissionMiddleware {
	return &RolePermissionMiddleware{
		auth: auth,
	}
}

func (m *RolePermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.TracerFromContext(r.Context()).Start(r.Context(), "RolePermissionMiddleware.Handle")
		defer span.End()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			logx.ErrorStack("empty metadata")
			http.Error(w, errUnauthorized, http.StatusUnauthorized)
			return
		}
		userIds := md.Get(userIdCtxKey)
		if len(userIds) == 0 {
			logx.ErrorStack("no user id")
			http.Error(w, errUnauthorized, http.StatusUnauthorized)
			return
		}
		userId, err := strconv.ParseUint(userIds[0], 10, 32)
		if err != nil {
			logx.ErrorStack(err)
			http.Error(w, errUnauthorized, http.StatusUnauthorized)
			return
		}
		permission, err := m.auth.CheckRolePermission(ctx, uint32(userId))
		if err != nil {
			if errors.Is(app.ErrAccessDenied, err) {
				http.Error(w, "permission denied", http.StatusForbidden)
				return
			}
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		if !permission {
			logx.Error("permission denied")
			http.Error(w, "permission denied", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
