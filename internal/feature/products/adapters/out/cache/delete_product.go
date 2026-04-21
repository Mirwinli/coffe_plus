package products_adapters_out_cache

import (
	"context"
	"fmt"

	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *CacheRepository) DeleteProduct(
	ctx context.Context,
	in products_ports_out.DeleteProductParams,
) error {
	product, ok := r.getProductFromCache(ctx, in.ID)
	if !ok {
		var (
			err error
		)

		repoGetProductResult, err := r.mainRepository.GetProduct(
			ctx,
			products_ports_out.NewGetProductParams(in.ID, false),
		)
		if err != nil {
			return fmt.Errorf("get task: %w", err)
		}
		product = repoGetProductResult.Product
	}

	r.invalidateProduct(ctx, &product.ID, &product.CategoryID)

	return r.mainRepository.DeleteProduct(ctx, in)
}
