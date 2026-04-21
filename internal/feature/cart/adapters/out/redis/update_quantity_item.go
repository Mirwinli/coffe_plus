package cart_adapters_out_redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (r *CartRepository) UpdateQuantityItem(
	ctx context.Context,
	in cart_ports_out.UpdateQuantityItemParams,
) (cart_ports_out.UpdateQuantityItemResult, error) {
	key := cartKey(in.CartID)

	if in.Quantity < 0 {
		cmd := r.pool.HExists(ctx, key, in.ProductID.String())
		if err := cmd.Err(); err != nil {
			return cart_ports_out.UpdateQuantityItemResult{}, fmt.Errorf(
				"check exists product in cart: %w", err,
			)
		}
		if !cmd.Val() {
			return cart_ports_out.UpdateQuantityItemResult{}, nil
		}
	}

	cmd := r.pool.HIncrBy(ctx, key, in.ProductID.String(), int64(in.Quantity))
	if err := cmd.Err(); err != nil {
		return cart_ports_out.UpdateQuantityItemResult{}, fmt.Errorf(
			"update quantity item HIncrBy: %w", err,
		)
	}

	val := cmd.Val()

	if val == 0 {
		if err := r.pool.HDel(ctx, key, in.ProductID.String()).Err(); err != nil {
			return cart_ports_out.UpdateQuantityItemResult{}, fmt.Errorf(
				"delete item from cache: %w", err,
			)
		}
	}

	cmdAll := r.pool.HGetAll(ctx, key)
	if err := cmdAll.Err(); err != nil {
		return cart_ports_out.UpdateQuantityItemResult{}, fmt.Errorf(
			"get all fields from cache: %w", err,
		)
	}

	values := cmdAll.Val()

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
	products, err := r.productsRepository.GetProductsByIDs(ctx, params)
	if err != nil {
		return cart_ports_out.UpdateQuantityItemResult{}, fmt.Errorf(
			"get products by ids: %w", err,
		)
	}

	createdAt, updatedAt, err := getCreatedUpdatedAt(values)
	if err != nil {
		return cart_ports_out.UpdateQuantityItemResult{}, fmt.Errorf(
			"get created updated At: %w", err,
		)
	}

	var totalPrice domain.Money

	var items []domain.Item
	for _, product := range products.Products {
		quantityStr := values[product.ID.String()]

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			return cart_ports_out.UpdateQuantityItemResult{}, fmt.Errorf(
				"convert quantity to integer: %w", err,
			)
		}

		item := domain.NewItem(
			product.ID,
			quantity,
			product.Price,
			product.Name,
			product.ImageURL,
		)

		linePrice := product.Price.Mul(decimal.NewFromInt(int64(quantity)))
		totalPrice = totalPrice.Add(linePrice)
		items = append(items, item)
	}

	cart := domain.NewCart(
		in.CartID,
		items,
		totalPrice,
		createdAt,
		updatedAt,
	)

	return cart_ports_out.NewUpdateQuantityItemResult(cart), nil
}
