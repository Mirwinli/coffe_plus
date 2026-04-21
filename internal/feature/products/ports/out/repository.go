package products_ports_out

import (
	"context"
)

type ProductRepository interface {
	SaveProduct(
		ctx context.Context,
		in SaveProductParams,
	) (SaveProductResult, error)
	GetProduct(
		ctx context.Context,
		in GetProductParams,
	) (GetProductResult, error)
	GetProducts(
		ctx context.Context,
		in GetProductsParams,
	) (GetProductsResult, error)
	UpdateProduct(
		ctx context.Context,
		in UpdateProductParams,
	) (UpdateProductResult, error)
	DeleteProduct(
		ctx context.Context,
		in DeleteProductParams,
	) error
	GetProductsByIDs(
		ctx context.Context,
		in GetProductsByIDsParams,
	) (GetProductsByIDsResult, error)
}
