package products_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetProductParams struct {
	ID uuid.UUID
}

func NewGetProductParams(id uuid.UUID) GetProductParams {
	return GetProductParams{ID: id}
}

type GetProductResult struct {
	Product domain.Product
}

func NewGetProductResult(product domain.Product) GetProductResult {
	return GetProductResult{Product: product}
}
