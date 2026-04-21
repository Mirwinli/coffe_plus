package products_service

import (
	"context"
	"fmt"

	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (s *ProductsService) DeleteProduct(
	ctx context.Context,
	in products_ports_in.DeleteProductParams,
) error {
	params := products_ports_out.NewGetProductParams(in.ID, false)
	product, err := s.ProductsRepository.GetProduct(ctx, params)
	if err != nil {
		return fmt.Errorf(
			"get product from repository: %w", err,
		)
	}

	if err = s.ImageUploader.Delete(ctx, product.Product.PublicID); err != nil {
		return fmt.Errorf(
			"delete image uploader: %w", err,
		)
	}

	delParams := products_ports_out.NewDeleteProductParams(in.ID)
	if err = s.ProductsRepository.DeleteProduct(ctx, delParams); err != nil {
		return fmt.Errorf(
			"delete product from repository: %w", err,
		)
	}

	return nil
}
