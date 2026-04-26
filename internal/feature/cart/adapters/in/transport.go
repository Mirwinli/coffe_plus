package cart_adapters_in

import (
	"net/http"

	core_http_middleware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
)

type CartHTTPHandler struct {
	cartService  cart_ports_in.CartService
	JWTConfig    core_http_jwt.Config
	controlToken core_http_middleware.AccessTokenBlackList
}

func NewCartHTTPHandler(
	cartService cart_ports_in.CartService,
	jwtConfig core_http_jwt.Config,
	controlToken core_http_middleware.AccessTokenBlackList,
) *CartHTTPHandler {
	return &CartHTTPHandler{
		cartService:  cartService,
		JWTConfig:    jwtConfig,
		controlToken: controlToken,
	}
}

func (h *CartHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/cart",
			Handler: h.AddProductInCart,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/cart",
			Handler: h.ListCart,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/cart",
			Handler: h.UpdateQuantityItem,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/cart",
			Handler: h.DeleteCart,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
	}
}
