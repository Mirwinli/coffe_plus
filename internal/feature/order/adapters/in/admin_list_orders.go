package order_adapters_in

import (
	"net/http"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
)

type AdminListOrdersResponse struct {
	Orders []domain.Order
}

// AdminListOrders godoc
// @Summary Виведенення замовлень
// @Description Виведеняня замовленнь за їх статусом
// @DEscription Admin only
// @Tags order
// @Security BearerAuth
// @Produce json
// @Param status query string false "Order status filter"
// @Success 200 {object} AdminListOrdersResponse "Замовлення"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /admin/order [get]
func (h *OrderHTTPHandler) AdminListOrders(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	status := getStatusQueryParams(r)

	in := order_ports_in.NewAdminListOrdersParams(status)
	getOrdersResult, err := h.orderService.AdminListOrders(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to list orders",
		)
		return
	}

	orders := getOrdersResult.Orders

	response := AdminListOrdersResponse{
		Orders: orders,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getStatusQueryParams(r *http.Request) *string {
	return core_http_request.GetStringQueryParams(r, "status")
}
