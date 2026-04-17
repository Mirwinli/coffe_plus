package auth_postgres

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) GetAndDeleteRefreshToken(
	ctx context.Context,
	in auth_ports_out.GetRefreshParams,
) (auth_ports_out.GetRefreshResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM coffe_plus.refresh_tokens
 			  WHERE token_hash = $1 AND device_name = $2
 			  RETURNING user_id,expires_at;`

	row := r.pool.QueryRow(ctx, query, in.RefreshToken, in.DeviceName)

	var refreshModel RefreshModel
	if err := row.Scan(
		&refreshModel.UserID,
		&refreshModel.ExpiresAt,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return auth_ports_out.GetRefreshResult{}, fmt.Errorf(
				"refresh token: %w",
				core_errors.ErrNotFound,
			)
		}
		return auth_ports_out.GetRefreshResult{}, fmt.Errorf(
			"scan error: %w", err,
		)
	}

	return auth_ports_out.NewGetRefreshResult(refreshModel.ExpiresAt, refreshModel.UserID), nil
}
