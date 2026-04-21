package products_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetProductParams struct {
	ID            uuid.UUID
	OnlyAvailable bool
}

func NewGetProductParams(id uuid.UUID, available bool) GetProductParams {
	return GetProductParams{
		ID:            id,
		OnlyAvailable: available,
	}
}

type GetProductResult struct {
	Product domain.Product
}

func NewGetProductResult(product domain.Product) GetProductResult {
	return GetProductResult{
		Product: product,
	}
}
