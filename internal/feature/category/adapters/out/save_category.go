package category_adapters_out

import (
	"context"
	"errors"
	"fmt"

	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (r *CategoryRepository) SaveCategory(
	ctx context.Context,
	in categories_ports_out.SaveCategoryParams,
) (categories_ports_out.SaveCategoryResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO coffe_plus.category (id,name)
			  VALUES ($1,$2)
			  RETURNING id,name,version
  		      `

	row := r.pool.QueryRow(ctx, query, in.Category.ID, in.Category.Name)

	var model CategoryModel
	if err := row.Scan(
		&model.ID,
		&model.Name,
		&model.Version,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesUnique) {
			return categories_ports_out.SaveCategoryResult{}, fmt.Errorf(
				"name must be unique; %w",
				err,
			)
		}
		return categories_ports_out.SaveCategoryResult{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	domain := domainFromModel(model)

	return categories_ports_out.NewSaveCategoryResult(domain), nil
}
