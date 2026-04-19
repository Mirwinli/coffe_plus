package core_http_middleware
 
import (
	"context"
	"fmt"
	"net/http"
 
	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
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
 
type AccessTokenBlackList interface {
	IsUserBlackListed(
		ctx context.Context,
		params auth_ports_out.IsBlackListedParams,
	) (bool, error)
}
 
func BlackListAccessToken(control AccessTokenBlackList) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
 
			id, ok := ctx.Value(core_contextKeys.JWTAccessIDCtxKey).(string)
			if !ok {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"jwt id invalid",
				)
				return
			}
 
			isBlacklisted, err := control.IsUserBlackListed(
				ctx,
				auth_ports_out.NewIsBlackListedParams(id),
			)
			if err != nil {
				responseHandler.ErrorResponse(
					err,
					"failed to check access token",
				)
				return
			}
 
			if isBlacklisted {
				responseHandler.ErrorResponse(
					core_errors.ErrUnauthorized,
					"access denied",
				)
				return
			}
 
			next.ServeHTTP(w, r)
		})
	}
}