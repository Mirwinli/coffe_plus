package order_service

import (
	"context"
	"fmt"

	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (s *OrderService) ListOrder(ctx context.Context, in order_ports_in.ListOrderParams) (order_ports_in.ListOrderResult, error) {
	params := order_ports_out.NewGetOrderParams(in.OrderID)

	getOrderResult, err := s.orderRepository.GetOrder(ctx, params)
	if err != nil {
		return order_ports_in.ListOrderResult{}, fmt.Errorf(
			"list order from repository: %w", err,
		)
	}

	return order_ports_in.NewListOrderResult(getOrderResult.Order), nil
}
