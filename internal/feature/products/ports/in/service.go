package products_ports_in

import "context"

type ProductService interface {
	CreateProduct(
		ctx context.Context,
		in CreateProductParams,
	) (CreateProductResult, error)
	GetProduct(
		ctx context.Context,
		in GetProductParams,
	) (GetProductResult, error)
	GetProducts(
		ctx context.Context,
		in GetProductsParams,
	) (GetProductsResult, error)
	PatchProduct(
		ctx context.Context,
		in PatchProductParams,
	) (PatchProductResult, error)
	DeleteProduct(
		ctx context.Context,
		in DeleteProductParams,
	) error
}
