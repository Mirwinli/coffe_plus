package products_adapters_out_cache

import (
	"context"

	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *CacheRepository) UpdateProduct(
	ctx context.Context,
	in products_ports_out.UpdateProductParams,
) (products_ports_out.UpdateProductResult, error) {
	patchedProduct, err := r.mainRepository.UpdateProduct(ctx, in)
	if err != nil {
		return products_ports_out.UpdateProductResult{}, err
	}

	product := patchedProduct.Product

	r.cacheProduct(ctx, product)
	r.invalidateProduct(ctx, &in.Product.ID, &in.Product.CategoryID)

	return patchedProduct, nil
}
