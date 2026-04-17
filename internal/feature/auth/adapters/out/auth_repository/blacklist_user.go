package auth_postgres

import (
	"context"
	"fmt"
	"time"

	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) BlackListUser(
	ctx context.Context,
	in auth_ports_out.BlackListParams,
) error {
	ttl := time.Until(in.ExpiresAt.Time)

	cmd := r.redisPool.Set(ctx, "blacklist:"+in.IDAccess.String(), "1", ttl)
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("set blacklist: %w", err)
	}

	return nil
}
