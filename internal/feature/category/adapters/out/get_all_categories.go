package category_adapters_out

import (
	"context"
	"fmt"

	categories_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/out"
)

func (r *CategoryRepository) GetAllCategories(
	ctx context.Context,
	params categories_ports_out.GetAllCategoriesParams,
) (categories_ports_out.GetAllCategoriesResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,name
			  FROM coffe_plus.category
			  LIMIT $1
			  OFFSET $2;
			  `

	rows, err := r.pool.Query(ctx, query, params.Limit, params.Offset)
	if err != nil {
		return categories_ports_out.GetAllCategoriesResult{}, fmt.Errorf(
			"select categories: %w", err,
		)
	}
	defer rows.Close()

	var models []CategoryModel
	for rows.Next() {
		var model CategoryModel

		if err = rows.Scan(
			&model.ID,
			&model.Version,
			&model.Name,
		); err != nil {
			return categories_ports_out.GetAllCategoriesResult{}, fmt.Errorf(
				"scan categories: %w", err,
			)
		}

		models = append(models, model)
	}
	if err = rows.Err(); err != nil {
		return categories_ports_out.GetAllCategoriesResult{}, fmt.Errorf(
			"next rows: %w", err,
		)
	}

	domains := domainsFromModels(models)

	return categories_ports_out.NewGetAllCategoriesResult(domains), nil
}
