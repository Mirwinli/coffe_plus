package category_adapters_out

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (r *CategoryRepository) UpdateCategory(
	ctx context.Context,
	in categories_ports_out.UpdateCategoryParams,
) (categories_ports_out.UpdateCategoryResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `UPDATE coffe_plus.category
			  SET	
			      name = $1,
			  	  version = version + 1
			  WHERE id = $2 AND version = $3
			  RETURNING id,name,version;
			  `
	row := r.pool.QueryRow(
		ctx,
		query,
		in.Category.Name,
		in.Category.ID,
		in.Category.Version,
	)

	var model CategoryModel
	if err := row.Scan(
		&model.ID,
		&model.Name,
		&model.Version,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return categories_ports_out.UpdateCategoryResult{}, fmt.Errorf(
				"task with id='%s' concurrently accessed: %w",
				in.Category.ID,
				core_errors.ErrConflict,
			)
		}
		return categories_ports_out.UpdateCategoryResult{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	domain := domainFromModel(model)
	return categories_ports_out.NewUpdateCategoryResult(domain), nil
}
