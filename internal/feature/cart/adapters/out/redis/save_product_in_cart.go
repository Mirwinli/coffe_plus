package cart_adapters_out_redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
	products_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/out"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (r *CartRepository) SaveProductInCart(
	ctx context.Context,
	in cart_ports_out.SaveProductInCartParams,
) (cart_ports_out.SaveProductInCartResult, error) {
	params := products_ports_out.NewGetProductParams(in.ProductID, true)
	getProductResult, err := r.productsRepository.GetProduct(ctx, params)
	if err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"get product from repository: %w", err,
		)
	}

	if !getProductResult.Product.IsAvailable {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"product isn't available: %w",
			core_errors.ErrProductIsntAvailable,
		)
	}

	key := cartKey(in.CartID)

	if err = r.pool.HIncrBy(ctx, key, getProductResult.Product.ID.String(), int64(in.Quantity)).Err(); err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"HIncrBy product: %w", err,
		)
	}

	now := time.Now().Unix()

	if err = r.pool.HSetNX(ctx, key, fieldCreatedAt, now).Err(); err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"HSetNX created_at: %w", err,
		)
	}
	if err = r.pool.HSet(ctx, key, fieldUpdatedAt, now).Err(); err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"HSetNX updated_at: %w", err,
		)
	}

	if err = r.pool.Expire(ctx, key, dayExpire).Err(); err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"set expire redis key: %w", err,
		)
	}

	cmd := r.pool.HGetAll(ctx, key)
	if err = cmd.Err(); err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"HGetAll: %w", err,
		)
	}

	values := cmd.Val()

	productIDs := []uuid.UUID{}
	for k, _ := range values {
		if k == fieldUpdatedAt || k == fieldCreatedAt {
			continue
		}

		if id, err := uuid.Parse(k); err == nil {
			productIDs = append(productIDs, id)
		}
	}

	idsParams := products_ports_out.NewGetProductsByIDsParams(productIDs)
	products, err := r.productsRepository.GetProductsByIDs(ctx, idsParams)
	if err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"get products by repository: %w", err,
		)
	}

	var items = []domain.Item{}
	var totalPrice domain.Money
	for _, product := range products.Products {
		quantityStr := values[product.ID.String()]

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
				"invalid quantity convert to integer: %w", err,
			)
		}

		linePrice := product.Price.Mul(decimal.NewFromInt(int64(quantity)))
		totalPrice = totalPrice.Add(linePrice)
		item := domain.NewItem(product.ID, quantity, product.Price, product.Name, product.ImageURL)
		items = append(items, item)
	}

	createdAt, updatedAt, err := getCreatedUpdatedAt(values)
	if err != nil {
		return cart_ports_out.SaveProductInCartResult{}, fmt.Errorf(
			"get created and updated at: %w", err,
		)
	}

	cart := domain.NewCart(
		in.CartID,
		items,
		totalPrice,
		createdAt,
		updatedAt,
	)

	return cart_ports_out.NewSaveProductInCartResult(cart), nil
}
