package cart_adapters_in

import (
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
	"github.com/google/uuid"
)

// DeleteCart godoc
// @Summary Видалення кошика
// @Description Видалення кошика та його вміст
// @Tags cart
// @Security BearerAuth
// @Success 204
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /cart [delete]
func (h *CartHTTPHandler) DeleteCart(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	cartID, ok := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrUnauthorized,
			"failed to get userID",
		)
		return
	}

	in := cart_ports_in.NewDeleteCartParams(cartID)
	if err := h.cartService.DeleteCart(ctx, in); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete cart",
		)
		return
	}

	responseHandler.NoContentResponse()
}
