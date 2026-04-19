package products_adapters_out_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *ProductsRepository) DeleteProduct(
	ctx context.Context,
	in products_ports_out.DeleteProductParams,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM coffe_plus.products WHERE id = $1`

	cmd, err := r.pool.Exec(ctx, query, in.ID)
	if err != nil {
		return fmt.Errorf("delete product exec: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf(
			"product not found: %w",
			core_errors.ErrNotFound,
		)
	}

	return nil
}
