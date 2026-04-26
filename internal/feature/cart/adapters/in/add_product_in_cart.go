package cart_adapters_in

import (
	"errors"
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
	"github.com/google/uuid"
)

type AddProductInCartRequest struct {
	ProductID uuid.UUID `json:"product_id" example:"ba930185-467f-4031-b1bd-abf4899dffer"`
	Quantity  int       `json:"quantity"   example:"12"`
}

type AddProductInCartResponse CartDTOResponse

// AddProductInCart
// @Summary Додавання продукта в кошик та його створення
// @Description Додавання продукта в кошик та одночасне створення кошика якщо він не був створений до цього часу
// @Security BearerAuth
// @Tags cart
// @Accept json
// @Produce json
// @Param request body AddProductInCartRequest true "AddProductInCart тіло запиту"
// @Success 201 {object} AddProductInCartResponse "Кошик"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /cart [post]
func (h *CartHTTPHandler) AddProductInCart(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request AddProductInCartRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode request",
		)
		return
	}

	cartID := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)

	in := cart_ports_in.NewAddProductInCartParams(request.ProductID, cartID, request.Quantity)
	newCart, err := h.cartService.AddProductInCart(ctx, in)
	if err != nil {
		if errors.Is(err, core_errors.ErrProductIsntAvailable) {
			responseHandler.ErrorResponse(
				err,
				"product is not available",
			)
			return
		}
		responseHandler.ErrorResponse(
			err,
			"failed to add product in cart",
		)
		return
	}

	response := AddProductInCartResponse{
		Cart: newCart.Cart,
	}

	responseHandler.JSONResponse(response, http.StatusCreated)
}
