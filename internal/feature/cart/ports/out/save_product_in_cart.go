package cart_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type SaveProductInCartParams struct {
	CartID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
}

func NewSaveProductInCartParams(productID uuid.UUID, cartID uuid.UUID, quantity int) SaveProductInCartParams {
	return SaveProductInCartParams{
		ProductID: productID,
		CartID:    cartID,
		Quantity:  quantity,
	}
}

type SaveProductInCartResult struct {
	Cart domain.Cart
}

func NewSaveProductInCartResult(cart domain.Cart) SaveProductInCartResult {
	return SaveProductInCartResult{
		Cart: cart,
	}
}
