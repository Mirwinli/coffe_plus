package order_service

import (
	"context"
	"fmt"

	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (s *OrderService) ListOrders(
	ctx context.Context,
	in order_ports_in.ListOrdersParams,
) (order_ports_in.ListOrdersResult, error) {
	params := order_ports_out.NewGetOrdersParams(in.UserID, in.Limit, in.Offset)

	getOrdersResult, err := s.orderRepository.GetOrders(ctx, params)
	if err != nil {
		return order_ports_in.ListOrdersResult{}, fmt.Errorf(
			"get orders from repository: %w", err,
		)
	}

	return order_ports_in.NewListOrdersResult(getOrdersResult.Orders), nil
}
