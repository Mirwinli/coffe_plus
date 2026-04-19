package products_adapters_in_products_transport_http

import (
	"net/http"

	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
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
}

func NewProductsHTTPHandler(productsService products_ports_in.ProductService) *ProductsHTTPHandler {
	return &ProductsHTTPHandler{
		productsService: productsService,
	}
}

func (h *ProductsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: h.CreateProduct,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products/{id}",
			Handler: h.GetProduct,
		},
		{
			Method:  http.MethodGet,
			Path:    "/products",
			Handler: h.GetProducts,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/products/{id}",
			Handler: h.PatchProduct,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/products/{id}",
			Handler: h.DeleteProduct,
		},
	}
}
