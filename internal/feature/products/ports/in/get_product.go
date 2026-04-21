package products_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetProductParams struct {
	ID            uuid.UUID
	OnlyAvailable bool
}

func NewGetProductParams(id uuid.UUID, onlyAvailable bool) GetProductParams {
	return GetProductParams{
		ID:            id,
		OnlyAvailable: onlyAvailable,
	}
}

type GetProductResult struct {
	Product domain.Product
}

func NewGetProductResult(product domain.Product) GetProductResult {
	return GetProductResult{Product: product}
}
