package core_http_middleware

import (
	"context"
	"net/http"
	"time"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

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

func ParseJWTToken(config core_http_jwt.Config) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			claims, err := getAndParseToken(r, config)
			if err != nil {
				responseHandler.ErrorResponse(
					err,
					"failed to parse JWT token",
				)
				return
			}

			ctxWithValue := context.WithValue(ctx, core_contextKeys.UserIDCtxKey, claims.UserID)
			ctxWithValue = context.WithValue(ctxWithValue, core_contextKeys.UserRoleCtxKey, claims.Role)
			ctxWithValue = context.WithValue(ctxWithValue, core_contextKeys.JWTAccessIDCtxKey, claims.ID)

			next.ServeHTTP(w, r.WithContext(ctxWithValue))
		})
	}
}

func CORS(allowedOriginsList []string) Middleware {
	allowedOrigins := make(map[string]struct{})
	for _, origin := range allowedOriginsList {
		allowedOrigins[origin] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			if requestID == "" {
				requestID = uuid.NewString()

				w.Header().Set(requestIDHeader, requestID)
				r.Header.Set(requestIDHeader, requestID)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
			)

			ctx := core_logger.ToContext(r.Context(), l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HTTP request got unexpected panic",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.String("method", r.Method),
				zap.Time("time", time.Now().UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Now().Sub(before)),
			)
		})
	}
}
