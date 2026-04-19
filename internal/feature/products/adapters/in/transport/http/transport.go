package products_adapters_in_products_transport_http

import (
	"net/http"

	core_http_middleware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
)

var (
	allowedFormatImage = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
)

type ProductsHTTPHandler struct {
	productsService products_ports_in.ProductService
	JWTConfig       core_http_jwt.Config
	controlToken    core_http_middleware.AccessTokenBlackList
}

func NewProductsHTTPHandler(
	productsService products_ports_in.ProductService,
	jwtConfig core_http_jwt.Config,
	controlToken core_http_middleware.AccessTokenBlackList,
) *ProductsHTTPHandler {
	return &ProductsHTTPHandler{
		productsService: productsService,
		JWTConfig:       jwtConfig,
		controlToken:    controlToken,
	}
}

func (h *ProductsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: h.CreateProduct,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/products/{id}",
			Handler: h.GetProduct,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: h.GetProducts,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/products/{id}",
			Handler: h.PatchProduct,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/products/{id}",
			Handler: h.DeleteProduct,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
	}
}
