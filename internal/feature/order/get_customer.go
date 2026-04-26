package order_service

import (
	"context"
	"fmt"

	order_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/in"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (s *OrderService) GetCustomer(ctx context.Context, in order_ports_in.GetCustomerParams) (order_ports_in.GetCustomerResult, error) {
	params := order_ports_out.NewGetCustomerParams(in.UserID)

	getCustomerResult, err := s.orderRepository.GetCustomer(ctx, params)
	if err != nil {
		return order_ports_in.GetCustomerResult{}, fmt.Errorf(
			"get customer from repository: %w", err,
		)
	}
	return order_ports_in.NewGetCustomerResult(getCustomerResult.User), nil
}
