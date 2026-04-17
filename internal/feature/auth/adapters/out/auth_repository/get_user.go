package auth_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) GetUser(
	ctx context.Context,
	in auth_ports_out.GetUserParams,
) (auth_ports_out.GetUserResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,password_hash,first_name,last_name,created_at,email,phone_number,role
			  FROM coffe_plus.users
			  WHERE id = $1;
			  `

	row := r.pool.QueryRow(ctx, query, in.UserID)

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
			return auth_ports_out.GetUserResult{}, fmt.Errorf(
				"user not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return auth_ports_out.GetUserResult{}, fmt.Errorf(
			"scan error: %w", err,
		)
	}

	user := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.PasswordHash,
		userModel.FirstName,
		userModel.LastName,
		userModel.Email,
		userModel.PhoneNumber,
		userModel.Role,
		userModel.CreatedAt,
	)

	return auth_ports_out.GetUserResult{
		User: user,
	}, nil
}
