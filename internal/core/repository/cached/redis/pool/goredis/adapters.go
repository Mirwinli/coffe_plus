package core_goredis_pool

import (
	"errors"

	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	"github.com/redis/go-redis/v9"
)

type StringCmd struct {
	*redis.StringCmd
}

func (c StringCmd) Bytes() ([]byte, error) {
	data, err := c.StringCmd.Bytes()
	if err != nil {
		return nil, mapError(err)
	}

	return data, nil
}

type StatusCmd struct {
	*redis.StatusCmd
}

type IntCmd struct {
	*redis.IntCmd
}

func mapError(err error) error {
	if errors.Is(err, redis.Nil) {
		return core_redis_pool.ErrNil
	}

	return err
}
