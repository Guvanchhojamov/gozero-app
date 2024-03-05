package middleware

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/repository"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
)

type RolePermissionMiddleware struct {
	authRepo repository.AuthRepository
}

func NewRolePermissionMiddleware(authRepo repository.AuthRepository) *RolePermissionMiddleware {
	return &RolePermissionMiddleware{
		authRepo: authRepo,
	}
}

func (m *RolePermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		md, ok := metadata.FromOutgoingContext(r.Context())
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
		permission, err := m.authRepo.CheckRolePermission(r.Context(), uint32(userId))
		if err != nil {
			if err.Error() == ErrAccessDenied.Error() {
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
