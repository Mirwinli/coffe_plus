package shop_ports_out

import "context"

type ShopRepository interface {
	ChangeShopStatus(
		ctx context.Context,
		in CangeShopStatusParams,
	) error
}
