package products_adapters_in_products_transport_http

import (
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
)

type GetProductResponse ProductDTOResponse

// GetProduct godoc
// @Summary Отримання продукту
// @Description Отримання продукту за його ID
// @Security BearerAuth
// @Tags product
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} GetProductResponse "Продукт"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Not found"
// @Router /products/{id} [get]
func (h *ProductsHTTPHandler) GetProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get id path value",
		)
		return
	}
	role := ctx.Value(core_contextKeys.UserRoleCtxKey)

	OnlyAvailable := true
	if role == "admin" {
		OnlyAvailable = false
	}

	in := products_ports_in.NewGetProductParams(id, OnlyAvailable)

	result, err := h.productsService.GetProduct(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get product",
		)
		return
	}

	response := GetProductResponse(productDTOFromDomain(result.Product))

	responseHandler.JSONResponse(response, http.StatusOK)
}
