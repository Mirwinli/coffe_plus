package shop_adapters_out_redis

import (
	"context"
	"fmt"

	shop_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/shop/ports/out"
)

const (
	ShopStatusKey = "shop:status"
)

func (r *ShopRepository) ChangeShopStatus(
	ctx context.Context,
	in shop_ports_out.CangeShopStatusParams,
) error {
	if err := r.pool.Set(ctx, ShopStatusKey, in.Status, 0).Err(); err != nil {
		return fmt.Errorf(
			"set shop status: %w", err,
		)
	}
	return nil
}
