package order_adapters_in

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
)

type ListOrderResponse OrderDTOResponse

// ListOrder godoc
// @Summary Отримання замовлення
// @Description Отримання замовлення по його ID
// @Tags order
// @Security BearerAuth
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} ListOrderResponse "Замовлення"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /order/{id} [get]
func (h *OrderHTTPHandler) ListOrder(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	orderID, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get order id path value",
		)
		return
	}

	in := order_ports_in.NewListOrderParams(orderID)
	listOrderResult, err := h.orderService.ListOrder(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to list order",
		)
		return
	}

	response := ListOrderResponse{
		Order: listOrderResult.Order,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
