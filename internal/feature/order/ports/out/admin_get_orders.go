package order_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type AdminGetOrdersParams struct {
	Status *string
}

func NewAdminGetOrdersParams(status *string) AdminGetOrdersParams {
	return AdminGetOrdersParams{
		Status: status,
	}
}

type AdminGetOrdersResult struct {
	Orders []domain.Order
}

func NewAdminGetOrdersResult(orders []domain.Order) AdminGetOrdersResult {
	return AdminGetOrdersResult{
		Orders: orders,
	}
}
