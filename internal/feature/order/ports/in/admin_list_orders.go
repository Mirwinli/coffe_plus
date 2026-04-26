package order_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
)

type AdminListOrdersParams struct {
	Status *string
}

func NewAdminListOrdersParams(
	status *string,
) AdminListOrdersParams {
	return AdminListOrdersParams{
		Status: status,
	}
}

type AdminListOrdersResult struct {
	Orders []domain.Order
}

func NewAdminListOrdersResult(orders []domain.Order) AdminListOrdersResult {
	return AdminListOrdersResult{
		Orders: orders,
	}
}
