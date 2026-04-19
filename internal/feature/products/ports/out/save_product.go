package products_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type SaveProductParams struct {
	Product domain.Product
}

func NewSaveProductParams(product domain.Product) SaveProductParams {
	return SaveProductParams{
		Product: product,
	}
}

type SaveProductResult struct {
	Product domain.Product
}

func NewSaveProductResult(product domain.Product) SaveProductResult {
	return SaveProductResult{
		Product: product,
	}
}
