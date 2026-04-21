package cart_adapters_out_redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_redis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (r *CartRepository) GetCart(
	ctx context.Context,
	in cart_ports_out.GetCartParams,
) (cart_ports_out.GetCartResult, error) {
	key := cartKey(in.ID)

	cmd := r.pool.HGetAll(ctx, key)
	if err := cmd.Err(); err != nil {
		if errors.Is(err, core_redis_pool.ErrNotFound) {
			return cart_ports_out.GetCartResult{}, fmt.Errorf(
				"cart not found: %w",
				core_errors.ErrNotFound,
			)
		}
		return cart_ports_out.GetCartResult{}, fmt.Errorf(
			"get product from cache: %w",
		)
	}

	values := cmd.Val()

	var productsID []uuid.UUID

	for k, _ := range values {
		if k == fieldUpdatedAt || k == fieldCreatedAt {
			continue
		}
		if productID, err := uuid.Parse(k); err == nil {
			productsID = append(productsID, productID)
		}
	}

	params := products_ports_out.NewGetProductsByIDsParams(productsID)
	getProductsResult, err := r.productsRepository.GetProductsByIDs(ctx, params)
	if err != nil {
		return cart_ports_out.GetCartResult{}, err
	}

	createdAt, updatedAt, err := getCreatedUpdatedAt(values)
	if err != nil {
		return cart_ports_out.GetCartResult{}, fmt.Errorf(
			"get created and updated at: %w", err,
		)
	}

	products := getProductsResult.Products

	var totalPrice domain.Money
	var items []domain.Item

	for _, product := range products {
		quantityStr := values[product.ID.String()]

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			return cart_ports_out.GetCartResult{}, fmt.Errorf(
				"convert quantity to int: %w",
			)
		}

		linePrice := product.Price.Mul(decimal.NewFromInt(int64(quantity)))
		totalPrice = totalPrice.Add(linePrice)

		item := domain.NewItem(
			product.ID,
			quantity,
			product.Price,
			product.Name,
			product.ImageURL,
		)

		items = append(items, item)
	}

	cart := domain.NewCart(
		in.ID,
		items,
		totalPrice,
		createdAt,
		updatedAt,
	)

	return cart_ports_out.NewGetCartResult(cart), nil
}
