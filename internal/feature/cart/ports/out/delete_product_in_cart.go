package cart_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type UpdateQuantityItemParams struct {
	CartID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
}

func NewUpdateQuantityItemParams(
	cartID uuid.UUID,
	productID uuid.UUID,
	quantity int,
) UpdateQuantityItemParams {
	return UpdateQuantityItemParams{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}
}

type UpdateQuantityItemResult struct {
	Cart domain.Cart
}

func NewUpdateQuantityItemResult(cart domain.Cart) UpdateQuantityItemResult {
	return UpdateQuantityItemResult{
		Cart: cart,
	}
}
