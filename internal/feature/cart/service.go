package cart_service

import (
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
)

type CartService struct {
	cartRepository cart_ports_out.CartRepository
}

func NewCartService(cart cart_ports_out.CartRepository) *CartService {
	return &CartService{
		cartRepository: cart,
	}
}
