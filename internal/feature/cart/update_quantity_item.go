package cart_service

import (
	"context"
	"fmt"

	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
)

func (s *CartService) UpdateQuantityItem(
	ctx context.Context,
	in cart_ports_in.UpdateQuantityItemParams,
) (cart_ports_in.UpdateQuantityItemResult, error) {
	params := cart_ports_out.NewUpdateQuantityItemParams(in.CartID, in.ProductID, in.Quantity)

	cart, err := s.cartRepository.UpdateQuantityItem(ctx, params)
	if err != nil {
		return cart_ports_in.UpdateQuantityItemResult{}, fmt.Errorf(
			"update quantity item: %w", err,
		)
	}

	return cart_ports_in.NewUpdateQuantityItemResult(cart.Cart), nil
}
