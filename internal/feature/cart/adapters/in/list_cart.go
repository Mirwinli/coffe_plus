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

type ListCartResponse CartDTOResponse

func (h *CartHTTPHandler) ListCart(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	cartID, ok := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrInvalidArgument,
			"invalid user id",
		)
		return
	}

	in := cart_ports_in.NewListCartParams(cartID)
	cart, err := h.cartService.ListCart(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to list cart",
		)
		return
	}

	response := ListCartResponse{
		Cart: cart.Cart,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
