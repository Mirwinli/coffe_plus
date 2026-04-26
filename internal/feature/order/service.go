package order_service

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

type OrderService struct {
	orderRepository order_ports_out.OrderRepository
	cartRepository  cart_ports_out.CartRepository
	orderNotifier   domain.OrderNotifier
}

func NewOrderService(
	orderRepository order_ports_out.OrderRepository,
	cartRepository cart_ports_out.CartRepository,
	orderNotifier domain.OrderNotifier,
) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		cartRepository:  cartRepository,
		orderNotifier:   orderNotifier,
	}
}
