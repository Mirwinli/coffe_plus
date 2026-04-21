package cart_adapters_out_redis

import (
	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

type CartRepository struct {
	pool               core_redis_pool.Pool
	productsRepository products_ports_out.ProductRepository
}

func NewCartRepository(
	redisPool core_redis_pool.Pool,
	productRepository products_ports_out.ProductRepository,
) *CartRepository {
	return &CartRepository{
		pool:               redisPool,
		productsRepository: productRepository,
	}
}
