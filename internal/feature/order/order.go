package order_service

import (
	"context"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	cart_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/cart/ports/out"
	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (s *OrderService) Order(
	ctx context.Context,
	in order_ports_in.OrderParams,
) (order_ports_in.OrderResult, error) {
	params := cart_ports_out.NewGetCartParams(in.OrderReceiver.CustomerID)

	getCartResult, err := s.cartRepository.GetCart(ctx, params)
	if err != nil {
		return order_ports_in.OrderResult{}, fmt.Errorf(
			"get cart from cache: %w", err,
		)
	}

	cart := getCartResult.Cart

	orderItems := domain.NewOrderItems(cart.Items)

	order := domain.NewOrderUnitialized(
		orderItems,
		cart.Price,
		in.OrderReceiver,
	)

	orderParams := order_ports_out.NewSaveOrderParams(order)
	saveOrderResult, err := s.orderRepository.SaveOrder(ctx, orderParams)
	if err != nil {
		return order_ports_in.OrderResult{}, fmt.Errorf(
			"save order from repository: %w", err,
		)
	}

	if err := s.orderNotifier.SendEmail(saveOrderResult.Order, saveOrderResult.Order.Status); err != nil {
		return order_ports_in.OrderResult{}, fmt.Errorf(
			"send email: %w", err,
		)
	}

	delParams := cart_ports_out.NewDeleteCartParams(in.OrderReceiver.CustomerID)
	if err := s.cartRepository.DeleteCart(ctx, delParams); err != nil {
		return order_ports_in.OrderResult{}, fmt.Errorf(
			"delete cart from cache: %w", err,
		)
	}

	return order_ports_in.NewOrderResult(saveOrderResult.Order), nil
}
