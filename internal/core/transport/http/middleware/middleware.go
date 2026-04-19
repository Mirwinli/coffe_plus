package core_http_middleware

import (
	"fmt"
	"net/http"
	"strings"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
)

const (
	AuthorizationHeaderKey = "Authorization"
	requestIDHeader        = "X-Request-ID"
)

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) == 0 {
		return h
	}

	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}

func getAndParseToken(
	r *http.Request,
	config core_http_jwt.Config,
) (*core_http_jwt.Claims, error) {
	authHeader := r.Header.Get(AuthorizationHeaderKey)
	if authHeader == "" {
		return &core_http_jwt.Claims{}, fmt.Errorf("no Authorization header found: %w",
			core_errors.ErrUnauthorized,
		)
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := core_http_jwt.ParseToken(accessToken, config)
	if err != nil {
		return &core_http_jwt.Claims{}, fmt.Errorf("parse JWT token: %w", err)
	}

	return claims, nil
}
