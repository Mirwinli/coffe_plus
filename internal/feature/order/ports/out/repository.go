package order_ports_out

import "context"

type OrderRepository interface {
	SaveOrder(
		ctx context.Context,
		in SaveOrderParams,
	) (SaveOrderResult, error)
	GetOrders(
		ctx context.Context,
		in GetOrdersParams,
	) (GetOrdersResult, error)
	AdminGetOrders(
		ctx context.Context,
		in AdminGetOrdersParams,
	) (AdminGetOrdersResult, error)
	GetOrder(
		ctx context.Context,
		in GetOrderParams,
	) (GetOrderResult, error)
	UpdateOrder(
		ctx context.Context,
		in UpdateOrderParams,
	) (UpdateOrderResult, error)
	GetCustomer(
		ctx context.Context,
		in GetCustomerParams,
	) (GetCustomerResult, error)
}
