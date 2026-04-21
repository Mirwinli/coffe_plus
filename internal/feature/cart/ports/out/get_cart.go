package cart_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetCartParams struct {
	ID uuid.UUID
}

func NewGetCartParams(id uuid.UUID) GetCartParams {
	return GetCartParams{
		ID: id,
	}
}

type GetCartResult struct {
	Cart domain.Cart
}

func NewGetCartResult(cart domain.Cart) GetCartResult {
	return GetCartResult{
		Cart: cart,
	}
}
