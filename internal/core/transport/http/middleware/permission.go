package core_http_middleware

import (
	"fmt"
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	"go.uber.org/zap"
)

func Admin() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			role, ok := ctx.Value(core_contextKeys.UserRoleCtxKey).(string)
			log.Debug("role from context", zap.String("role", role), zap.Bool("ok", ok))

			if !ok || role != "admin" {
				responseHandler.ErrorResponse(
					fmt.Errorf(
						"permission denied: %w",
						core_errors.ErrForbidden,
					),
					"Permission denied",
				)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
