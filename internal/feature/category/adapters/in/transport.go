package category_adapters_in

import (
	"net/http"

	core_http_middleware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	category_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

type CategoryHTTPHandler struct {
	categoryService category_ports_in.CategoryService
	JWTConfig       core_http_jwt.Config
	controlToken    core_http_middleware.AccessTokenBlackList
}

func NewCategoryHTTPHandler(
	categoryService category_ports_in.CategoryService,
	jwtConfig core_http_jwt.Config,
	controlToken core_http_middleware.AccessTokenBlackList,
) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{
		categoryService: categoryService,
		JWTConfig:       jwtConfig,
		controlToken:    controlToken,
	}
}

func (h *CategoryHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/category",
			Handler: h.GetAllCategories,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/category/{id}",
			Handler: h.GetCategory,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodPost,
			Path:    "/category",
			Handler: h.CreateCategory,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/category/{id}",
			Handler: h.PatchCategory,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/category/{id}",
			Handler: h.DeleteCategory,
			Middleware: []core_http_middleware.Middleware{
				core_http_middleware.ParseJWTToken(h.JWTConfig),
				core_http_middleware.Admin(),
				core_http_middleware.BlackListAccessToken(h.controlToken),
			},
		},
	}
}
