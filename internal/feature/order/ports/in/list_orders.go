package order_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type ListOrdersParams struct {
	UserID uuid.UUID
	Limit  *int
	Offset *int
}

func NewListOrdersParams(
	userID uuid.UUID,
	limit *int,
	offset *int,
) ListOrdersParams {
	return ListOrdersParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}
}

type ListOrdersResult struct {
	Orders []domain.Order
}

func NewListOrdersResult(orders []domain.Order) ListOrdersResult {
	return ListOrdersResult{
		Orders: orders,
	}
}
