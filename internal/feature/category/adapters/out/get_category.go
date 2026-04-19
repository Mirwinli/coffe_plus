package category_adapters_out

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (r *CategoryRepository) GetCategory(
	ctx context.Context,
	in categories_ports_out.GetCategoryParams,
) (categories_ports_out.GetCategoryResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,name
			  FROM coffe_plus.category
 			  WHERE id = $1;`

	row := r.pool.QueryRow(ctx, query, in.ID)

	var model CategoryModel
	if err := row.Scan(
		&model.ID,
		&model.Version,
		&model.Name,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return categories_ports_out.GetCategoryResult{}, fmt.Errorf(
				"category not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return categories_ports_out.GetCategoryResult{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	domain := domainFromModel(model)

	return categories_ports_out.NewGetCategoryResult(domain), nil
}
