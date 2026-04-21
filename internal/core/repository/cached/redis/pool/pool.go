package core_redis_pool

import (
	"context"
	"time"
)

type Pool interface {
	Get(ctx context.Context, key string) StringCmd
	Set(ctx context.Context, key string, value any, ttl time.Duration) StatusCmd
	Del(ctx context.Context, key ...string) IntCmd
	HGet(ctx context.Context, key string, field string) StringCmd
	HGetAll(ctx context.Context, key string) MapStringStringCmd
	HSet(ctx context.Context, key string, values ...any) IntCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) IntCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) BoolCmd
	HDel(ctx context.Context, key string, fields ...string) IntCmd
	HExists(ctx context.Context, key, field string) BoolCmd
	Expire(ctx context.Context, key string, expiration time.Duration) BoolCmd
	Close() error

	TTL() time.Duration
}

type StringCmd interface {
	Bytes() ([]byte, error)
}

type StatusCmd interface {
	Err() error
}

type IntCmd interface {
	Err() error
	Val() int64
}

type BoolCmd interface {
	Val() bool
	Err() error
}

type MapStringStringCmd interface {
	Val() map[string]string
	Err() error
}
