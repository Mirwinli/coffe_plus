package products_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetProductsByIDsParams struct {
	ID []uuid.UUID
}

func NewGetProductsByIDsParams(ids []uuid.UUID) GetProductsByIDsParams {
	return GetProductsByIDsParams{
		ID: ids,
	}
}

type GetProductsByIDsResult struct {
	Products []domain.Product
}

func NewGetProductsByIDsResult(products []domain.Product) GetProductsByIDsResult {
	return GetProductsByIDsResult{
		Products: products,
	}
}
