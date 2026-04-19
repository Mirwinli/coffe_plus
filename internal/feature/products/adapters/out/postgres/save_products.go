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

func (r *ProductsRepository) SaveProduct(
	ctx context.Context,
	in products_ports_out.SaveProductParams,
) (products_ports_out.SaveProductResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO coffe_plus.products (id,name,description,price,is_available,public_id,image_url,category_id)
			  VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
			  RETURNING id,version,name,description,price,is_available,public_id,image_url,category_id;
				`
	product := in.Product
	row := r.pool.QueryRow(
		ctx,
		query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.IsAvaible,
		product.PublicID,
		product.ImageURL,
		product.CategoryID,
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
		if errors.Is(err, core_postgres_pool.ErrViolatesUnique) {
			return products_ports_out.SaveProductResult{}, fmt.Errorf(
				"duplicate product: %w",
				core_errors.ErrUniqueViolation,
			)
		}
		return products_ports_out.SaveProductResult{}, fmt.Errorf(
			"scan error: %w",
			err,
		)
	}

	domainProduct := domain.NewProduct(
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
	result := products_ports_out.NewSaveProductResult(domainProduct)

	return result, nil
}
