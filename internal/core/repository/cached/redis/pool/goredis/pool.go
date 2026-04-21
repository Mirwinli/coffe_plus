package core_goredis_pool

import (
	"context"
	"fmt"
	"time"

	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	"github.com/redis/go-redis/v9"
)

type Pool struct {
	client *redis.Client
	ttl    time.Duration
}

func NewPool(ctx context.Context, config Config) (*Pool, error) {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		Username: "",
	}

	client := redis.NewClient(options)

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping : %w", err)
	}

	return &Pool{
		client,
		config.TTL,
	}, nil
}

func (p *Pool) Get(ctx context.Context, key string) core_redis_pool.StringCmd {
	cmd := p.client.Get(ctx, key)

	return goredisStringCmd{cmd}
}

func (p *Pool) Set(ctx context.Context, key string, value any, ttl time.Duration) core_redis_pool.StatusCmd {
	cmd := p.client.Set(ctx, key, value, ttl)

	return goredisStatusCmd{cmd}
}

func (p *Pool) Del(ctx context.Context, key ...string) core_redis_pool.IntCmd {
	cmd := p.client.Del(ctx, key...)

	return goredisIntCmd{cmd}
}

func (p *Pool) HGet(ctx context.Context, key string, field string) core_redis_pool.StringCmd {
	cmd := p.client.HGet(ctx, key, field)

	return goredisStringCmd{cmd}
}

func (p *Pool) HSet(ctx context.Context, key string, values ...any) core_redis_pool.IntCmd {
	cmd := p.client.HSet(ctx, key, values...)

	return goredisIntCmd{cmd}
}

func (p *Pool) HIncrBy(ctx context.Context, key, field string, incr int64) core_redis_pool.IntCmd {
	cmd := p.client.HIncrBy(ctx, key, field, incr)

	return goredisIntCmd{cmd}
}

func (p *Pool) HGetAll(ctx context.Context, key string) core_redis_pool.MapStringStringCmd {
	cmd := p.client.HGetAll(ctx, key)

	return goredisMapStringCmd{cmd}
}

func (p *Pool) HSetNX(ctx context.Context, key, field string, value interface{}) core_redis_pool.BoolCmd {
	cmd := p.client.HSetNX(ctx, key, field, value)

	return goredisBoolCmd{cmd}
}

func (p *Pool) Expire(ctx context.Context, key string, expiration time.Duration) core_redis_pool.BoolCmd {
	cmd := p.client.Expire(ctx, key, expiration)

	return goredisBoolCmd{cmd}
}

func (p *Pool) HDel(ctx context.Context, key string, fields ...string) core_redis_pool.IntCmd {
	cmd := p.client.HDel(ctx, key, fields...)

	return goredisIntCmd{cmd}
}

func (p *Pool) HExists(ctx context.Context, key, field string) core_redis_pool.BoolCmd {
	cmd := p.client.HExists(ctx, key, field)

	return goredisBoolCmd{cmd}
}

func (p *Pool) Close() error {
	if err := p.client.Close(); err != nil {
		return err
	}

	return nil
}

func (p *Pool) TTL() time.Duration {
	return p.ttl
}
