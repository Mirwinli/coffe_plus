package category_adapters_out

import (
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
)

type CategoryRepository struct {
	pool core_postgres_pool.Pool
}

func NewCategoryRepository(pool core_postgres_pool.Pool) *CategoryRepository {
	return &CategoryRepository{pool: pool}
}
