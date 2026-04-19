package products_adapters_out_cache

import (
	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

type CacheRepository struct {
	pool           core_redis_pool.Pool
	mainRepository products_ports_out.ProductRepository
}

func NewCacheRepository(
	pool core_redis_pool.Pool,
	mainRepository products_ports_out.ProductRepository,
) *CacheRepository {
	return &CacheRepository{
		pool:           pool,
		mainRepository: mainRepository,
	}
}
