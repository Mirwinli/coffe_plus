package auth_postgres

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) LoginUser(
	ctx context.Context,
	in auth_ports_out.LoginUserAuthParams,
) (auth_ports_out.LoginUserAuthResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,first_name,last_name,created_at,email,phone_number,password_hash,role
			  FROM coffe_plus.users
			  WHERE email = $1;
			`

	row := r.pool.QueryRow(ctx, query, in.Email)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.CreatedAt,
		&userModel.Email,
		&userModel.PhoneNumber,
		&userModel.PasswordHash,
		&userModel.Role,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return auth_ports_out.LoginUserAuthResult{}, fmt.Errorf(
				"user not found:%w",
				core_errors.ErrNotFound,
			)
		}
		return auth_ports_out.LoginUserAuthResult{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	return auth_ports_out.NewLoginUserAuthResult(
		userModel.ID,
		userModel.Version,
		userModel.PasswordHash,
		userModel.FirstName,
		userModel.LastName,
		userModel.Email,
		userModel.PhoneNumber,
		userModel.Role,
	), nil
}
