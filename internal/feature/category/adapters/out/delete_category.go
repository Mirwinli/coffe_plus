package category_adapters_out

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (r *CategoryRepository) DeleteCategory(
	ctx context.Context,
	in categories_ports_out.DeleteCategoryParams,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM coffe_plus.category
			  WHERE id = $1;`

	cmd, err := r.pool.Exec(ctx, query, in.ID)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrForeignKeyViolation) {
			return core_errors.ErrForeignKeyViolation
		}
		return fmt.Errorf("exec delete category: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf(
			"category not found; %w",
			core_errors.ErrNotFound,
		)
	}

	return nil
}
