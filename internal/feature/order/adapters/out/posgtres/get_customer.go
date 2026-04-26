package order_adapters_out_posgtres

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	auth_postgres "github.com/Mirwinli/coffe_plus/internal/feature/auth/adapters/out/auth_repository"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (r *OrderRepository) GetCustomer(
	ctx context.Context,
	in order_ports_out.GetCustomerParams,
) (order_ports_out.GetCustomerResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,first_name,last_name,email,phone_number,role,created_at
			  FROM coffe_plus.users
			 WHERE id = $1; 
			  `

	var userModel auth_postgres.UserModel
	if err := r.pool.QueryRow(ctx, query, in.UserID).Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.Email,
		&userModel.PhoneNumber,
		&userModel.Role,
		&userModel.CreatedAt,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return order_ports_out.GetCustomerResult{}, fmt.Errorf(
				"customer not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return order_ports_out.GetCustomerResult{}, fmt.Errorf(
			"scan error: %w", err,
		)
	}

	user := auth_postgres.DomainFromModel(userModel)

	return order_ports_out.NewGetCustomerResult(user), nil
}
