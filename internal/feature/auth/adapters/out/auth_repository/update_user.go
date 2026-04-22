package auth_postgres

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) UpdateUser(ctx context.Context, in auth_ports_out.UpdateUserParams) (auth_ports_out.UpdateUserResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `UPDATE coffe_plus.users
			  SET
				  first_name = $1,
				  last_name = $2,
				  password_hash = $3,
				  phone_number = $4,
				  version = version + 1
			  WHERE id = $5 AND version = $6
			  RETURNING id,version,password_hash,first_name,last_name,created_at,email,phone_number,role
				  `

	user := in.User
	row := r.pool.QueryRow(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Password,
		user.PhoneNumber,
		user.ID,
		user.Version,
	)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.PasswordHash,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.CreatedAt,
		&userModel.Email,
		&userModel.PhoneNumber,
		&userModel.Role,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return auth_ports_out.UpdateUserResult{}, fmt.Errorf(
				"user not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return auth_ports_out.UpdateUserResult{}, fmt.Errorf(
			"scan error: %w", err,
		)
	}

	user = domainFromModel(userModel)

	return auth_ports_out.NewUpdateUserResult(user), nil
}
