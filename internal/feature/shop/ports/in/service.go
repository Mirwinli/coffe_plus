package shop_ports_in

import "context"

type ShopService interface {
	CangeShopStatus(
		ctx context.Context,
		in ChangeShopStatusParams,
	) error
}
