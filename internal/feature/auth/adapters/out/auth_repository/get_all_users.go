package auth_postgres

import (
	"context"
	"fmt"

	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (r *AuthRepository) GetAllUsers(
	ctx context.Context,
	in auth_ports_out.GetAllUsersParams,
) (auth_ports_out.GetAllUsersResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,first_name,last_name,created_at,email,phone_number,role
			  FROM coffe_plus.users	
			  LIMIT $1
			  OFFSET $2; 
			  `

	rows, err := r.pool.Query(ctx, query, in.Limit, in.Offset)
	if err != nil {
		return auth_ports_out.GetAllUsersResult{}, fmt.Errorf(
			"select error: %w", err,
		)
	}

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel

		if err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FirstName,
			&userModel.LastName,
			&userModel.CreatedAt,
			&userModel.Email,
			&userModel.PhoneNumber,
			&userModel.Role,
		); err != nil {
			return auth_ports_out.GetAllUsersResult{}, fmt.Errorf(
				"scan error: %w", err,
			)
		}
		userModels = append(userModels, userModel)
	}

	if err := rows.Err(); err != nil {
		return auth_ports_out.GetAllUsersResult{}, fmt.Errorf(
			"next error: %w", err,
		)
	}

	return auth_ports_out.NewGetAllUsersResult(DomainsFromModels(userModels)), nil
}
