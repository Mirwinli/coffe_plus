package adapters_in_auth_transport_http

import (
	"net/http"

	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
)

type AuthHTTPHandler struct {
	authService auth_ports_in.AuthService
	JWTConfig   core_http_jwt.Config
}

func NewAuthHTTPHandler(
	authService auth_ports_in.AuthService,
	config core_http_jwt.Config,
) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
		JWTConfig:   config,
	}
}

func (h *AuthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/register",
			Handler: h.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/refresh",
			Handler: h.Refresh,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: h.Login,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/auth/logout",
			Handler: h.Logout,
		},
	}
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}

	return r.RemoteAddr
}
