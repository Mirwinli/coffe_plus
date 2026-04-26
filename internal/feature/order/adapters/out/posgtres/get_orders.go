package order_adapters_out_posgtres

import (
	"context"
	"encoding/json"
	"fmt"

	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (r *OrderRepository) GetOrders(
	ctx context.Context,
	in order_ports_out.GetOrdersParams,
) (order_ports_out.GetOrdersResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT o.id,o.version,o.customer_id,o.status,o.total_price,o.created_at,u.id,u.phone_number,u.email,u.first_name,u.last_name,
			  json_agg(json_build_object(
			  'id',i.id,
			  'order_id',i.order_id,
			  'product_id',i.product_id,
			  'product_name',i.product_name,
			  'image_url',i.image_url,
			  'quantity',i.quantity,
			  'price_at_time',i.price_at_time
			  )) AS items
	FROM coffe_plus.orders o
	LEFT JOIN coffe_plus.order_items i ON o.id = i.order_id
	JOIN coffe_plus.users u ON o.customer_id = u.id
	WHERE o.customer_id = $1
	GROUP BY o.id, u.id, u.phone_number, u.email, u.first_name, u.last_name
	ORDER BY o.created_at DESC
	LIMIT $2 OFFSET $3;
	`

	rows, err := r.pool.Query(ctx, query, in.UserID, in.Limit, in.Offset)
	if err != nil {
		return order_ports_out.GetOrdersResult{}, fmt.Errorf(
			"select orders: %w", err,
		)
	}
	defer rows.Close()

	var orders []OrderModel
	for rows.Next() {
		var orderModel OrderModel
		var itemsRaw []byte

		if err := rows.Scan(
			&orderModel.ID,
			&orderModel.Version,
			&orderModel.OrderReceiver.CustomerID,
			&orderModel.Status,
			&orderModel.Price,
			&orderModel.CreatedAt,
			&orderModel.OrderReceiver.CustomerID,
			&orderModel.OrderReceiver.PhoneNumber,
			&orderModel.OrderReceiver.Email,
			&orderModel.OrderReceiver.FirstName,
			&orderModel.OrderReceiver.LastName,
			&itemsRaw,
		); err != nil {
			return order_ports_out.GetOrdersResult{}, fmt.Errorf(
				"scan error: %w", err,
			)
		}
		if len(itemsRaw) != 0 {
			if err := json.Unmarshal(itemsRaw, &orderModel.Items); err != nil {
				return order_ports_out.GetOrdersResult{}, fmt.Errorf(
					"unmarshal items: %w", err,
				)
			}
		}

		orders = append(orders, orderModel)
	}

	domains := orderModelsToDomains(orders)

	return order_ports_out.NewGetOrdersResult(domains), nil
}
