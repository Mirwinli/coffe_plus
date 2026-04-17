package auth_postgres

import (
	"context"
	"fmt"

	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) SaveRefreshToken(
	ctx context.Context,
	in auth_ports_out.SaveRefreshTokenAuthParams,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO coffe_plus.refresh_tokens
			  (token_hash,user_id,device_name,ip_address,expires_at,created_at)
			  VALUES 
			  ($1,$2,$3,$4,$5,$6)
			  ON CONFLICT (user_id,device_name)
			  DO UPDATE SET
			     token_hash = EXCLUDED.token_hash,
			     ip_address = EXCLUDED.ip_address,
			     expires_at = EXCLUDED.expires_at,
			     created_at = EXCLUDED.created_at
`

	cmd, err := r.pool.Exec(
		ctx,
		query,
		in.TokenHash,
		in.UserID,
		in.DeviceName,
		in.IpAddress,
		in.ExpiresAt,
		in.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("exec refresh token: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("exec refresh token: no rows affected")
	}

	return nil
}
