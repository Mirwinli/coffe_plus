package order_adapters_in

import (
	"net/http"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
)

type PatchOrderRequest struct {
	Status string `json:"status",validate:"required" example:"completed"`
}

type PatchOrderResponse struct {
	Order domain.Order `json:"order"`
}

// PatchOrder
// @Summary  Зміна статусу замовлення
// @Description Зміна статусу замовлення
// @Description Admin Only
// @Security BearerAuth
// @Tags order
// @Accept json
// @Produce json
// @Param request body PatchOrderRequest true "PatchOrder тіло запиту"
// @Params id path string true "Order ID"
// @Success 200 {object} PatchOrderResponse "Оновлене замовлення"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /order/{id} [patch]
func (h *OrderHTTPHandler) PatchOrder(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	orderID, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `orderID` path value",
		)
		return
	}

	var request PatchOrderRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	in := order_ports_in.NewUpdateOrderParams(request.Status, orderID)
	updateOrderResult, err := h.orderService.UpdateOrder(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch order",
		)
		return
	}

	response := PatchOrderResponse{Order: updateOrderResult.Order}

	responseHandler.JSONResponse(response, http.StatusOK)
}
