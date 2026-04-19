package products_adapters_in_products_transport_http

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
)

func (h *ProductsHTTPHandler) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get id from path",
		)
		return
	}

	in := products_ports_in.NewDeleteProductParams(id)
	if err = h.productsService.DeleteProduct(ctx, in); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete product",
		)
		return
	}

	responseHandler.NoContentResponse()
}
