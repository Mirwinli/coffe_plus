package order_service

import (
	"context"
	"fmt"

	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (s *OrderService) UpdateOrder(
	ctx context.Context,
	in order_ports_in.UpdateOrderParams,
) (order_ports_in.UpdateOrderResult, error) {

	orderParams := order_ports_out.NewGetOrderParams(in.OrderID)
	getOrderResult, err := s.orderRepository.GetOrder(ctx, orderParams)
	if err != nil {
		return order_ports_in.UpdateOrderResult{}, fmt.Errorf(
			"get order from repository: %w", err,
		)
	}

	if err := getOrderResult.Order.ValidateStatusTransitons(in.Status); err != nil {
		return order_ports_in.UpdateOrderResult{}, fmt.Errorf(
			"status validation error: %w", err,
		)
	}

	getOrderResult.Order.Status = in.Status

	params := order_ports_out.NewUpdateOrderParams(getOrderResult.Order)

	updateOrderResult, err := s.orderRepository.UpdateOrder(ctx, params)
	if err != nil {
		return order_ports_in.UpdateOrderResult{}, fmt.Errorf(
			"update order from repository: %w", err,
		)
	}

	order := updateOrderResult.Order

	if err := s.orderNotifier.SendEmail(order, order.Status); err != nil {
		return order_ports_in.UpdateOrderResult{}, fmt.Errorf(
			"send email: %w", err,
		)
	}

	return order_ports_in.NewUpdateOrderResult(order), nil
}
