package cart_service

import (
	"context"
	"fmt"

	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
)

func (s *CartService) ListCart(
	ctx context.Context,
	in cart_ports_in.ListCartParams,
) (cart_ports_in.ListCartResult, error) {
	params := cart_ports_out.NewGetCartParams(in.ID)

	getCartResult, err := s.cartRepository.GetCart(ctx, params)
	if err != nil {
		return cart_ports_in.ListCartResult{}, fmt.Errorf(
			"get cart from repository: %w", err,
		)
	}

	return cart_ports_in.NewListCartResult(getCartResult.Cart), err
}
