package categories_ports_out

import "github.com/google/uuid"

type CountProductsParams struct {
	CategoryID uuid.UUID
}

func NewCountProductsParams(id uuid.UUID) CountProductsParams {
	return CountProductsParams{
		CategoryID: id,
	}
}

type CountProductsResult struct {
	Count int
}

func NewCountProductsResult(count int) CountProductsResult {
	return CountProductsResult{
		Count: count,
	}
}
