package products_adapters_out_cache

import (
	"context"

	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *CacheRepository) GetProductsByIDs(
	ctx context.Context,
	in products_ports_out.GetProductsByIDsParams,
) (products_ports_out.GetProductsByIDsResult, error) {
	products, err := r.mainRepository.GetProductsByIDs(ctx, in)
	if err != nil {
		return products_ports_out.GetProductsByIDsResult{}, err
	}

	return products, nil
}
