package auth_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) LogoutUser(ctx context.Context, in auth_ports_out.LogoutUserAuthParams) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM coffe_plus.refresh_tokens
			  WHERE token_hash = $1 AND device_name = $2
			  `

	cmd, err := r.pool.Exec(ctx, query, in.RefreshToken, in.DeviceName)
	if err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("refresh token not found: %w", core_errors.ErrNotFound)
	}

	return nil
}
