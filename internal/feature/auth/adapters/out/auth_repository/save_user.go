package auth_postgres

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) SaveUser(
	ctx context.Context,
	params auth_ports_out.SaveUserAuthParams,
) (auth_ports_out.SaveUserAuthResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO coffe_plus.users (id,password_hash,first_name,last_name,email,phone_number,created_at)
			  VALUES ($1,$2,$3,$4,$5,$6,$7)
			  RETURNING id,version,first_name,last_name,email,phone_number,role,created_at`

	user := params.User
	row := r.pool.QueryRow(
		ctx,
		query,
		user.ID,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.CreatedAt,
	)

	var userModel UserModel

	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.Email,
		&userModel.PhoneNumber,
		&userModel.Role,
		&userModel.CreatedAt,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesUnique) {
			return auth_ports_out.SaveUserAuthResult{}, fmt.Errorf(
				"%v user with id=%s: %w",
				err,
				user.ID,
				core_errors.ErrConflict,
			)
		}
		return auth_ports_out.SaveUserAuthResult{}, fmt.Errorf("scan error: %w", err)
	}

	return auth_ports_out.NewSaveUserAuthResult(
		userModel.ID,
		userModel.Version,
		userModel.FirstName,
		userModel.LastName,
		userModel.Email,
		userModel.PhoneNumber,
		userModel.Role,
		userModel.CreatedAt,
	), nil
}
