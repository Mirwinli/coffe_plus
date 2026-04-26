package order_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
)

type UpdateOrderParams struct {
	Order domain.Order
}

func NewUpdateOrderParams(
	order domain.Order,
) UpdateOrderParams {
	return UpdateOrderParams{
		Order: order,
	}
}

type UpdateOrderResult struct {
	Order domain.Order
}

func NewUpdateOrderResult(
	order domain.Order,
) UpdateOrderResult {
	return UpdateOrderResult{
		Order: order,
	}
}
