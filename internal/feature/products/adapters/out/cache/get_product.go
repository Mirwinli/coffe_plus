package products_adapters_out_cache

import (
	"context"

	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *CacheRepository) GetProduct(
	ctx context.Context,
	in products_ports_out.GetProductParams,
) (products_ports_out.GetProductResult, error) {
	if product, ok := r.getProductFromCache(ctx, in.ID); ok {
		result := products_ports_out.NewGetProductResult(product)
		return result, nil
	}

	params := products_ports_out.NewGetProductParams(in.ID)

	product, err := r.mainRepository.GetProduct(ctx, params)
	if err != nil {
		return products_ports_out.GetProductResult{}, err
	}

	r.cacheProduct(ctx, product.Product)

	return product, nil
}
