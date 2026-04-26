package order_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type UpdateOrderParams struct {
	OrderID uuid.UUID
	Status  string
}

func NewUpdateOrderParams(
	status string,
	orderID uuid.UUID,
) UpdateOrderParams {
	return UpdateOrderParams{
		Status:  status,
		OrderID: orderID,
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
