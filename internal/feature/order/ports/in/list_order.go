package order_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type ListOrderParams struct {
	OrderID uuid.UUID
}

func NewListOrderParams(orderID uuid.UUID) ListOrderParams {
	return ListOrderParams{
		OrderID: orderID,
	}
}

type ListOrderResult struct {
	Order domain.Order
}

func NewListOrderResult(order domain.Order) ListOrderResult {
	return ListOrderResult{
		Order: order,
	}
}
