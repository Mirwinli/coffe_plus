package products_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetProductsParams struct {
	Limit      *int
	Offset     *int
	CategoryID *uuid.UUID
}

func NewGetProductsParams(
	limit *int,
	offset *int,
	categoryID *uuid.UUID,
) GetProductsParams {
	return GetProductsParams{
		Limit:      limit,
		Offset:     offset,
		CategoryID: categoryID,
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
