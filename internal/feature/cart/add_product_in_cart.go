package cart_service

import (
	"context"
	"fmt"

	cart_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/in"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
)

func (s *CartService) AddProductInCart(
	ctx context.Context,
	in cart_ports_in.AddProductInCartParams,
) (cart_ports_in.AddProductInCartResult, error) {
	params := cart_ports_out.NewSaveProductInCartParams(in.ProductID, in.CartID, in.Quantity)

	cart, err := s.cartRepository.SaveProductInCart(ctx, params)
	if err != nil {
		return cart_ports_in.AddProductInCartResult{}, fmt.Errorf(
			"add product in cart by repository: %w", err,
		)
	}

	return cart_ports_in.NewAddProductInCartResult(cart.Cart), nil
}
