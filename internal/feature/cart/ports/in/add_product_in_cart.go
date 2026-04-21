package cart_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type AddProductInCartParams struct {
	CartID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
}

func NewAddProductInCartParams(productID uuid.UUID, cartID uuid.UUID, quantity int) AddProductInCartParams {
	return AddProductInCartParams{
		ProductID: productID,
		CartID:    cartID,
		Quantity:  quantity,
	}
}

type AddProductInCartResult struct {
	Cart domain.Cart
}

func NewAddProductInCartResult(cart domain.Cart) AddProductInCartResult {
	return AddProductInCartResult{
		Cart: cart,
	}
}
