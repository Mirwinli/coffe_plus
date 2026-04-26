package order_service

import (
	"context"
	"fmt"

	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (s *OrderService) AdminListOrders(
	ctx context.Context,
	in order_ports_in.AdminListOrdersParams,
) (order_ports_in.AdminListOrdersResult, error) {
	params := order_ports_out.NewAdminGetOrdersParams(in.Status)

	getOrdersResult, err := s.orderRepository.AdminGetOrders(ctx, params)
	if err != nil {
		return order_ports_in.AdminListOrdersResult{}, fmt.Errorf(
			"list admin orders: %w", err,
		)
	}

	return order_ports_in.NewAdminListOrdersResult(getOrdersResult.Orders), nil
}
