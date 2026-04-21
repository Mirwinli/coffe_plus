package cart_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type ListCartParams struct {
	ID uuid.UUID
}

func NewListCartParams(id uuid.UUID) ListCartParams {
	return ListCartParams{
		ID: id,
	}
}

type ListCartResult struct {
	Cart domain.Cart
}

func NewListCartResult(cart domain.Cart) ListCartResult {
	return ListCartResult{
		Cart: cart,
	}
}
