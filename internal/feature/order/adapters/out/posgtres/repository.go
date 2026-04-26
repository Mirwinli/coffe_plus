package order_adapters_out_posgtres

import (
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
)

type OrderRepository struct {
	pool core_postgres_pool.Pool
}

func NewOrderRepository(pool core_postgres_pool.Pool) *OrderRepository {
	return &OrderRepository{pool: pool}
}
