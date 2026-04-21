package products_adapters_out_postgres

import (
	"context"
	"fmt"

	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *ProductsRepository) GetProductsByIDs(
	ctx context.Context,
	in products_ports_out.GetProductsByIDsParams,
) (products_ports_out.GetProductsByIDsResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,name,description,price,is_available,category_id,public_id,image_url
			  FROM coffe_plus.products
			  WHERE id = ANY($1) AND is_available = TRUE;
			  `

	rows, err := r.pool.Query(ctx, query, in.ID)
	if err != nil {
		return products_ports_out.GetProductsByIDsResult{}, fmt.Errorf(
			"select products: %w", err,
		)
	}
	defer rows.Close()

	var products []ProductModel
	for rows.Next() {
		var product ProductModel

		if err = rows.Scan(
			&product.ID,
			&product.Version,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.IsAvaible,
			&product.CategoryID,
			&product.PublicID,
			&product.ImageURL,
		); err != nil {
			return products_ports_out.GetProductsByIDsResult{}, fmt.Errorf(
				"scan error: %w", err,
			)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products_ports_out.GetProductsByIDsResult{}, fmt.Errorf(
			"next rows: %w", err,
		)
	}

	return products_ports_out.NewGetProductsByIDsResult(modelsToDomains(products)), nil
}
