package order_adapters_in

import (
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type OrderResponse struct {
	Order domain.Order `json:"order"`
}

// Order godoc
// @Summary Зробити замовлення
// @Description Кинути запит на замовлення
// @Description Запит може бути відхиленим або прийнятим,надсилання йього запиту не гарантує його виконання
// @Description Дані для замовлення берутся напряму з кошика
// @Tags order
// @Security BearerAuth
// @Produce json
// @Success 201 {object} OrderResponse "Замовлення"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 404 {object} core_http_response.ErrorResponse "Cart not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /order [post]
func (h *OrderHTTPHandler) Order(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrInvalidArgument,
			"user id not found",
		)
		return
	}

	log.Debug("getting customer for user id", zap.String("user_id", userID.String()))

	getIn := order_ports_in.NewGetCustomerParams(userID)
	getCustomerResult, err := h.orderService.GetCustomer(ctx, getIn)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get customer",
		)
		return
	}

	customer := getCustomerResult.User

	orderReceiver := domain.NewOrderReceiver(
		customer.ID,
		customer.PhoneNumber,
		customer.Email,
		customer.FirstName,
		customer.LastName,
	)

	in := order_ports_in.NewOrderParams(orderReceiver)
	order, err := h.orderService.Order(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to order",
		)
		return
	}

	response := OrderResponse{
		Order: order.Order,
	}

	responseHandler.JSONResponse(response, http.StatusCreated)
}
