package products_adapters_out_postgres

import (
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
)

type ProductsRepository struct {
	pool core_postgres_pool.Pool
}

func NewRepository(pool core_postgres_pool.Pool) *ProductsRepository {
	return &ProductsRepository{pool: pool}
}
