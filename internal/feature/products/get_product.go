package products_service

import (
	"context"
	"fmt"

	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (s *ProductsService) GetProduct(
	ctx context.Context,
	in products_ports_in.GetProductParams,
) (products_ports_in.GetProductResult, error) {
	params := products_ports_out.NewGetProductParams(in.ID, in.OnlyAvailable)

	result, err := s.ProductsRepository.GetProduct(ctx, params)
	if err != nil {
		return products_ports_in.GetProductResult{}, fmt.Errorf(
			"get product from repository: %w", err,
		)
	}

	return products_ports_in.NewGetProductResult(result.Product), err
}
