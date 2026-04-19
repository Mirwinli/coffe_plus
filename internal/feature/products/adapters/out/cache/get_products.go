package products_adapters_out_cache

import (
	"context"
	"errors"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
	"go.uber.org/zap"
)

func (r *CacheRepository) GetProducts(
	ctx context.Context,
	in products_ports_out.GetProductsParams,
) (products_ports_out.GetProductsResult, error) {
	log := core_logger.FromContext(ctx)

	key := productsListKey(in.CategoryID)
	field := productsListField(in.Limit, in.Offset)

	bytes, err := r.pool.HGet(ctx, key, field).Bytes()
	if err != nil {
		if !errors.Is(err, core_redis_pool.ErrNotFound) {
			log.Error("hget product list", zap.Error(err))
		}
	} else {
		var productListModel ProductListModel
		if err = productListModel.Deserialize(bytes); err != nil {
			log.Error("deserialize product list", zap.Error(err))
		} else {
			products := modelToDomains(productListModel)

			return products_ports_out.NewGetProductsResult(products), nil
		}
	}

	repoGetProductsResult, err := r.mainRepository.GetProducts(ctx, in)
	if err != nil {
		return products_ports_out.GetProductsResult{}, err
	}

	productListModel := domainsToModel(repoGetProductsResult.Products)
	bytes, err = productListModel.Serialize()
	if err != nil {
		log.Error("serialize product list", zap.Error(err))
	} else {
		if err = r.pool.HSet(ctx, key, field, bytes).Err(); err != nil {
			log.Error("hset product list", zap.Error(err))
		}
	}
	return repoGetProductsResult, nil
}
