package core_http_server

import (
	"net/http"

	core_http_middlware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_http_middlware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_http_middlware.ChainMiddleware(r.Handler, r.Middleware...)
}
