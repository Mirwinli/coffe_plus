package products_service

import (
	"context"
	"fmt"

	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (s *ProductsService) PatchProduct(
	ctx context.Context,
	in products_ports_in.PatchProductParams,
) (products_ports_in.PatchProductResult, error) {
	getParams := products_ports_out.NewGetProductParams(in.ID, false)

	result, err := s.ProductsRepository.GetProduct(ctx, getParams)
	if err != nil {
		return products_ports_in.PatchProductResult{}, fmt.Errorf(
			"get product from repository: %w", err,
		)
	}

	product := result.Product

	err = product.ApplyPatch(in.Patch)
	if err != nil {
		return products_ports_in.PatchProductResult{}, fmt.Errorf(
			"apply patch: %w", err,
		)
	}

	patchParams := products_ports_out.NewUpdateProductParams(product)

	patchedProduct, err := s.ProductsRepository.UpdateProduct(ctx, patchParams)
	if err != nil {
		return products_ports_in.PatchProductResult{}, fmt.Errorf(
			"update product: %w", err,
		)
	}

	return products_ports_in.NewPatchProductResult(patchedProduct.Product), nil
}
