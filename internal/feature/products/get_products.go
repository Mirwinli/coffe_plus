package products_service

import (
	"context"
	"fmt"

	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (s *ProductsService) GetProducts(
	ctx context.Context,
	in products_ports_in.GetProductsParams,
) (products_ports_in.GetProductsResult, error) {
	params := products_ports_out.NewGetProductsParams(in.Limit, in.Offset, in.CategoryID, in.OnlyAvailable)

	products, err := s.ProductsRepository.GetProducts(ctx, params)
	if err != nil {
		return products_ports_in.GetProductsResult{}, fmt.Errorf(
			"get products from repository: %w", err,
		)
	}

	return products_ports_in.NewGetProductsResult(products.Products), nil
}
