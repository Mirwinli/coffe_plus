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

func (r *ProductsRepository) UpdateProduct(
	ctx context.Context,
	in products_ports_out.UpdateProductParams,
) (products_ports_out.UpdateProductResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `UPDATE coffe_plus.products 
			  SET 
			      name = $1,
			      description = $2,
			      price = $3,
			      is_available = $4,
			      category_id = $5,
			      version = version + 1
			   	WHERE id = $6 AND version = $7

			  RETURNING id, version,name ,description,price,is_available,public_id,image_url,category_id
`
	product := in.Product

	row := r.pool.QueryRow(
		ctx,
		query,
		product.Name,
		product.Description,
		product.Price,
		product.IsAvaible,
		product.CategoryID,
		product.ID,
		product.Version,
	)

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
			return products_ports_out.UpdateProductResult{}, fmt.Errorf(
				"task with id='%s' concurrently accessed: %w",
				product.ID,
				core_errors.ErrConflict,
			)
		}
		return products_ports_out.UpdateProductResult{}, fmt.Errorf(
			"scan error: %w", err,
		)
	}

	result := domain.NewProduct(
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

	return products_ports_out.NewUpdateProductResult(result), nil
}
