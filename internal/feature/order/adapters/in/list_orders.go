package order_adapters_in

import (
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	"github.com/google/uuid"
)

type ListOrdersResponse []OrderDTOResponse

// ListOrder godoc
// @Summary Виведення всіх замовлень
// @Description Виведення всіх замовленнь корисутвача який кинув запит
// @Tags order
// @Security BearerAuth
// @Produce json
// @Success 200 {object} ListOrdersResponse "Замовлення"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /order [get]
func (h *OrderHTTPHandler) ListOrders(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrInvalidArgument,
			"failed to get userID",
		)
		return
	}

	offset, limit, err := core_http_request.GetOffsetLimitQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get limit/offset query params",
		)
		return
	}

	in := order_ports_in.NewListOrdersParams(userID, limit, offset)
	getOrdersResult, err := h.orderService.ListOrders(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get orders",
		)
		return
	}

	orders := getOrdersResult.Orders

	response := ListOrdersResponse(orderDTOsFromDomains(orders))

	responseHandler.JSONResponse(response, http.StatusOK)
}
