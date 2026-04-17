package auth_postgres

import (
	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
)

type AuthRepository struct {
	pool      core_postgres_pool.Pool
	redisPool core_redis_pool.Pool
}

func NewAuthRepository(
	pool core_postgres_pool.Pool,
	redisPool core_redis_pool.Pool,
) *AuthRepository {
	return &AuthRepository{
		pool:      pool,
		redisPool: redisPool,
	}
}
