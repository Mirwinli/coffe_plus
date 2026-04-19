package category_adapters_in

import (
	"net/http"

	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	category_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

type CategoryHTTPHandler struct {
	categoryService category_ports_in.CategoryService
}

func NewCategoryHTTPHandler(categoryService category_ports_in.CategoryService) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/category",
			Handler: h.GetAllCategories,
		},
		{
			Method:  http.MethodGet,
			Path:    "/category/{id}",
			Handler: h.GetCategory,
		},
		{
			Method:  http.MethodPost,
			Path:    "/category",
			Handler: h.CreateCategory,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/category/{id}",
			Handler: h.PatchCategory,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/category/{id}",
			Handler: h.DeleteCategory,
		},
	}
}
