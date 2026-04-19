package products_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type UpdateProductParams struct {
	Product domain.Product
}

func NewUpdateProductParams(product domain.Product) UpdateProductParams {
	return UpdateProductParams{
		Product: product,
	}
}

type UpdateProductResult struct {
	Product domain.Product
}

func NewUpdateProductResult(product domain.Product) UpdateProductResult {
	return UpdateProductResult{
		Product: product,
	}
}
