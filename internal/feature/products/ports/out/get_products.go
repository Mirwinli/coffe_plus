package products_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetProductsParams struct {
	Limit         *int
	Offset        *int
	CategoryID    *uuid.UUID
	OnlyAvailable bool
}

func NewGetProductsParams(
	limit *int,
	offset *int,
	categoryID *uuid.UUID,
	onlyAvailable bool,
) GetProductsParams {
	return GetProductsParams{
		Limit:         limit,
		Offset:        offset,
		CategoryID:    categoryID,
		OnlyAvailable: onlyAvailable,
	}
}

type GetProductsResult struct {
	Products []domain.Product
}

func NewGetProductsResult(products []domain.Product) GetProductsResult {
	return GetProductsResult{
		Products: products,
	}
}
