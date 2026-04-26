package order_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetOrderParams struct {
	OrderID uuid.UUID
}

func NewGetOrderParams(orderID uuid.UUID) GetOrderParams {
	return GetOrderParams{
		OrderID: orderID,
	}
}

type GetOrderResult struct {
	Order domain.Order
}

func NewGetOrderResult(
	order domain.Order,
) GetOrderResult {
	return GetOrderResult{
		Order: order,
	}
}
