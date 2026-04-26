package order_ports_in

import "context"

type OrderService interface {
	Order(
		ctx context.Context,
		in OrderParams,
	) (OrderResult, error)
	ListOrders(
		ctx context.Context,
		in ListOrdersParams,
	) (ListOrdersResult, error)
	AdminListOrders(
		ctx context.Context,
		in AdminListOrdersParams,
	) (AdminListOrdersResult, error)
	ListOrder(
		ctx context.Context,
		in ListOrderParams,
	) (ListOrderResult, error)
	GetCustomer(
		ctx context.Context,
		in GetCustomerParams,
	) (GetCustomerResult, error)
	UpdateOrder(
		ctx context.Context,
		in UpdateOrderParams,
	) (UpdateOrderResult, error)
}
