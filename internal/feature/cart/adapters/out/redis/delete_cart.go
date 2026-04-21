package cart_adapters_out_redis

import (
	"context"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
)

func (r *CartRepository) DeleteCart(
	ctx context.Context,
	in cart_ports_out.DeleteCartParams,
) error {
	key := cartKey(in.ID)
	cmd := r.pool.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return fmt.Errorf("delete cart from repository: %w", err)
	}

	if cmd.Val() == 0 {
		return fmt.Errorf("cart not found in cache:%w", core_errors.ErrNotFound)
	}

	return nil
}
