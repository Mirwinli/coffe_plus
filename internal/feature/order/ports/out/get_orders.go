package order_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetOrdersParams struct {
	UserID uuid.UUID
	Limit  *int
	Offset *int
}

func NewGetOrdersParams(
	userID uuid.UUID,
	limit *int,
	offset *int,
) GetOrdersParams {
	return GetOrdersParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}
}

type GetOrdersResult struct {
	Orders []domain.Order
}

func NewGetOrdersResult(orders []domain.Order) GetOrdersResult {
	return GetOrdersResult{
		Orders: orders,
	}
}
