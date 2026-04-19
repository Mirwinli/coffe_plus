package products_adapters_out_cache

import (
	"context"
	"errors"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *CacheRepository) getProductFromCache(
	ctx context.Context,
	id uuid.UUID,
) (domain.Product, bool) {
	log := core_logger.FromContext(ctx)

	key := productKey(id)

	bytes, err := r.pool.Get(ctx, key).Bytes()
	if err != nil {
		if !errors.Is(err, core_redis_pool.ErrNotFound) {
			log.Error("read from cache", zap.Error(err))
		}
		return domain.Product{}, false
	}

	var productModel ProductModel
	if err = productModel.Deserialize(bytes); err != nil {
		log.Error("deserialized caches product", zap.Error(err))

		return domain.Product{}, false
	}

	productDomain := modelToDomain(productModel)

	return productDomain, true
}

func (r *CacheRepository) cacheProduct(
	ctx context.Context,
	product domain.Product,
) {
	log := core_logger.FromContext(ctx)

	productModel := domainToModel(product)

	bytes, err := productModel.Serialize()
	if err != nil {
		log.Error("serialize product", zap.Error(err))
	} else {
		if err = r.pool.Set(
			ctx,
			productKey(productModel.ID),
			bytes,
			r.pool.TTL(),
		).Err(); err != nil {
			log.Error("set product in cache", zap.Error(err))
		}
	}
}

func (r *CacheRepository) invalidateProduct(
	ctx context.Context,
	productID *uuid.UUID,
	categoryID *uuid.UUID,
) {
	log := core_logger.FromContext(ctx)

	invalidateKey := []string{
		productsListKey(nil),
	}

	if productID != nil {
		invalidateKey = append(invalidateKey, productKey(*productID))
	}

	if categoryID != nil {
		invalidateKey = append(invalidateKey, productKey(*categoryID))
	}

	if err := r.pool.Del(ctx, invalidateKey...).Err(); err != nil {
		log.Error("invalidate cached product list", zap.Error(err))
	}
}
