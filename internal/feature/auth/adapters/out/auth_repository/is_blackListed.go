package auth_postgres

import (
	"context"
	"errors"
	"fmt"

	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) IsUserBlackListed(
	ctx context.Context,
	params auth_ports_out.IsBlackListedParams,
) (bool, error) {

	cmd := r.redisPool.Get(ctx, "blacklist:"+params.IDAccess.String())
	if _, err := cmd.Bytes(); err != nil {
		if errors.Is(err, core_redis_pool.ErrNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("get access token redis: %w", err)
	}

	return true, nil
}
