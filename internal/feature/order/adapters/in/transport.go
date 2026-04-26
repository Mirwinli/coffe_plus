package order_adapters_in

import (
	"net/http"

	core_http_middleware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
)

type OrderHTTPHandler struct {
	orderService order_ports_in.OrderService
	JWTConfig    core_http_jwt.Config
	controlToken core_http_middleware.AccessTokenBlackList
}

func NewOrderHTTPHandler(
	orderService order_ports_in.OrderService,
	JWTConfig core_http_jwt.Config,
	controlToken core_http_middleware.AccessTokenBlackList,
) *OrderHTTPHandler {
	return &OrderHTTPHandler{
		orderService: orderService,
		JWTConfig:    JWTConfig,
		controlToken: controlToken,
	}
}

func (h *OrderHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Handler: h.Order,
			Path:    "/order",
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodGet,
			Handler: h.ListOrder,
			Path:    "/order/{id}",
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodGet,
			Handler: h.ListOrders,
			Path:    "/order",
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodPatch,
			Handler: h.PatchOrder,
			Path:    "/order/{id}",
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodGet,
			Handler: h.AdminListOrders,
			Path:    "/admin/order",
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
	}
}
