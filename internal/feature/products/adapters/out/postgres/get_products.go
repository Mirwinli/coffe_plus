package products_adapters_out_postgres

import (
	"context"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
)

func (r *ProductsRepository) GetProducts(
	ctx context.Context,
	in products_ports_out.GetProductsParams,
) (products_ports_out.GetProductsResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT id,version,name,description,price,is_available,public_id,image_url,category_id
			  FROM coffe_plus.products
			  %s    
			  ORDER BY id ASC
		      LIMIT $1
			  OFFSET $2;`

	where := ""
	args := []any{in.Limit, in.Offset}

	if in.CategoryID != nil {
		where = "WHERE category_id = $3"
		args = append(args, in.CategoryID)
	}

	finalQuery := fmt.Sprintf(query, where)

	rows, err := r.pool.Query(ctx, finalQuery, args...)
	if err != nil {
		return products_ports_out.GetProductsResult{}, fmt.Errorf(
			"select products: %w", err,
		)
	}
	defer rows.Close()

	var productModels []ProductModel
	for rows.Next() {
		var productModel ProductModel

		if err = rows.Scan(
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
			return products_ports_out.GetProductsResult{}, fmt.Errorf(
				"scan products: %w", err,
			)
		}
		productModels = append(productModels, productModel)
	}

	if err = rows.Err(); err != nil {
		return products_ports_out.GetProductsResult{}, fmt.Errorf(
			"next rows: %w", err,
		)
	}

	return products_ports_out.NewGetProductsResult(modelsToDomains(productModels)), nil
}

func modelsToDomains(models []ProductModel) []domain.Product {
	domains := make([]domain.Product, len(models))

	for i, model := range models {
		domains[i] = modelToDomain(model)
	}

	return domains
}

func modelToDomain(model ProductModel) domain.Product {
	return domain.NewProduct(
		model.ID,
		model.Version,
		model.Name,
		model.Description,
		model.Price,
		model.IsAvaible,
		model.CategoryID,
		model.ImageURL,
		model.PublicID,
	)
}
