package products_adapters_out_cache

import (
	"context"

	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *CacheRepository) SaveProduct(
	ctx context.Context,
	in products_ports_out.SaveProductParams,
) (products_ports_out.SaveProductResult, error) {
	repoSaveProductResult, err := r.mainRepository.SaveProduct(ctx, in)
	if err != nil {
		return products_ports_out.SaveProductResult{}, err
	}

	product := repoSaveProductResult.Product

	r.cacheProduct(ctx, product)

	r.invalidateProduct(ctx, nil, nil)

	return repoSaveProductResult, nil
}
