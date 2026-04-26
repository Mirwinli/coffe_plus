package order_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type SaveOrderParams struct {
	Order domain.Order
}

func NewSaveOrderParams(order domain.Order) SaveOrderParams {
	return SaveOrderParams{
		Order: order,
	}
}

type SaveOrderResult struct {
	Order domain.Order
}

func NewSaveOrderResult(order domain.Order) SaveOrderResult {
	return SaveOrderResult{
		Order: order,
	}
}
