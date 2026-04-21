package cart_ports_in

import "context"

type CartService interface {
	AddProductInCart(
		ctx context.Context,
		in AddProductInCartParams,
	) (AddProductInCartResult, error)
	ListCart(
		ctx context.Context,
		in ListCartParams,
	) (ListCartResult, error)
	UpdateQuantityItem(
		ctx context.Context,
		in UpdateQuantityItemParams,
	) (UpdateQuantityItemResult, error)
	DeleteCart(
		ctx context.Context,
		in DeleteCartParams,
	) error
}
