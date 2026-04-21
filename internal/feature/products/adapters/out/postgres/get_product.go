package products_adapters_out_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *ProductsRepository) GetProduct(
	ctx context.Context,
	in products_ports_out.GetProductParams,
) (products_ports_out.GetProductResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,name,description,price,is_available,public_id,image_url,category_id
			  FROM coffe_plus.products
			  WHERE id = $1`

	if in.OnlyAvailable {
		query += " AND is_available = TRUE;"
	}

	row := r.pool.QueryRow(ctx, query, in.ID)

	var productModel ProductModel
	if err := row.Scan(
		&productModel.ID,
		&productModel.Version,
		&productModel.Name,
		&productModel.Description,
		&productModel.Price,
		&productModel.IsAvaible,
		&productModel.PublicID,
		&productModel.ImageURL,
		&productModel.CategoryID,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return products_ports_out.GetProductResult{}, fmt.Errorf(
				"product not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return products_ports_out.GetProductResult{}, fmt.Errorf(
			"scan error: %w", err,
		)
	}

	product := domain.NewProduct(
		productModel.ID,
		productModel.Version,
		productModel.Name,
		productModel.Description,
		productModel.Price,
		productModel.IsAvaible,
		productModel.CategoryID,
		productModel.ImageURL,
		productModel.PublicID,
	)

	return products_ports_out.NewGetProductResult(product), nil
}
