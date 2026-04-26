package cart_adapters_in

import (
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
	"github.com/google/uuid"
)

type UpdateQuantityItemRequest struct {
	Quantity int `json:"quantity" example:"-1"`
}

type UpdateQuantityItemResponse struct {
	Cart domain.Cart `json:"cart"`
}

// UpdateQuantityItem godoc
// @Summary Зміна кількості придмету в кошиці
// @Description Зміна кількості продукту реалізована у вигляді додавання/віднімання
// @Tags cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} UpdateQuantityItemResponse "Оновлений кошик"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 409 {object core_http_response.ErrorResponse "Conflict"
// @Failure 401 {object} core_http_response.ErrorResponse "Unathorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /cart [patch]
func (h *CartHTTPHandler) UpdateQuantityItem(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	cartID, ok := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrUnauthorized,
			"failed to get user id",
		)
		return
	}

	productID, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `productID` path value",
		)
		return
	}

	var request UpdateQuantityItemRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	in := cart_ports_in.NewUpdateQuantityItemParams(cartID, productID, request.Quantity)
	cart, err := h.cartService.UpdateQuantityItem(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to update quantity item",
		)
		return
	}

	response := UpdateQuantityItemResponse{
		Cart: cart.Cart,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
