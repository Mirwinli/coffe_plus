package products_service

import (
	"context"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (s *ProductsService) CreateProduct(
	ctx context.Context,
	in products_ports_in.CreateProductParams,
) (products_ports_in.CreateProductResult, error) {

	imageURL, publicID, err := s.ImageUploader.Upload(ctx, in.ImageFile)
	if err != nil {
		return products_ports_in.CreateProductResult{}, fmt.Errorf(
			"upload file:%w",
			err,
		)
	}

	product := domain.NewProductUninitialized(
		in.Name,
		in.Description,
		in.Price,
		in.IsAvailable,
		in.CategoryID,
		imageURL,
		publicID,
	)

	if err = product.Validate(); err != nil {
		return products_ports_in.CreateProductResult{}, fmt.Errorf(
			"validation product:%w",
			err,
		)
	}

	params := products_ports_out.NewSaveProductParams(product)
	result, err := s.ProductsRepository.SaveProduct(ctx, params)
	if err != nil {
		return products_ports_in.CreateProductResult{}, fmt.Errorf(
			"save product from repository:%w",
			err,
		)
	}

	return products_ports_in.NewCreateProductResult(result.Product), nil
}
