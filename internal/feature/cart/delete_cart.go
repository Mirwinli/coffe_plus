package cart_service

import (
	"context"
	"fmt"

	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
)

func (s *CartService) DeleteCart(
	ctx context.Context,
	in cart_ports_in.DeleteCartParams,
) error {
	params := cart_ports_out.NewDeleteCartParams(in.ID)

	if err := s.cartRepository.DeleteCart(ctx, params); err != nil {
		return fmt.Errorf("delete cart from repository: %w", err)
	}

	return nil
}
