package shop_adapters_out_redis

import core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"

type ShopRepository struct {
	pool core_redis_pool.Pool
}

func NewShopRepository(pool core_redis_pool.Pool) ShopRepository {
	return ShopRepository{
		pool: pool,
	}
}
