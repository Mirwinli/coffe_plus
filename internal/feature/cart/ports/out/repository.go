package cart_ports_out

import "context"

type CartRepository interface {
	SaveProductInCart(
		ctx context.Context,
		in SaveProductInCartParams,
	) (SaveProductInCartResult, error)
	GetCart(
		ctx context.Context,
		in GetCartParams,
	) (GetCartResult, error)
	UpdateQuantityItem(
		ctx context.Context,
		in UpdateQuantityItemParams,
	) (UpdateQuantityItemResult, error)
	DeleteCart(
		ctx context.Context,
		in DeleteCartParams,
	) error
}
