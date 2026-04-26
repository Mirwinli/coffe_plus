package order_ports_in

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type OrderParams struct {
	OrderReceiver domain.OrderReceiver
}

func NewOrderParams(orderReceiver domain.OrderReceiver) OrderParams {
	return OrderParams{
		OrderReceiver: orderReceiver,
	}
}

type OrderResult struct {
	Order domain.Order
}

func NewOrderResult(order domain.Order) OrderResult {
	return OrderResult{
		Order: order,
	}
}
