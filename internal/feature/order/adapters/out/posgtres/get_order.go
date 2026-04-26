package order_adapters_out_posgtres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (r *OrderRepository) GetOrder(
	ctx context.Context,
	in order_ports_out.GetOrderParams,
) (order_ports_out.GetOrderResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT o.id,o.version,o.customer_id,o.status,o.total_price,o.created_at,u.first_name,u.last_name,u.email,u.phone_number,
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
			  WHERE o.id = $1
			  GROUP BY o.id,u.id; 
			  `

	row := r.pool.QueryRow(ctx, query, in.OrderID)

	var orderModel OrderModel
	var rawItems []byte
	if err := row.Scan(
		&orderModel.ID,
		&orderModel.Version,
		&orderModel.OrderReceiver.CustomerID,
		&orderModel.Status,
		&orderModel.Price,
		&orderModel.CreatedAt,
		&orderModel.OrderReceiver.FirstName,
		&orderModel.OrderReceiver.LastName,
		&orderModel.OrderReceiver.Email,
		&orderModel.OrderReceiver.PhoneNumber,
		&rawItems,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return order_ports_out.GetOrderResult{}, fmt.Errorf(
				"order not found: %w", core_errors.ErrNotFound,
			)
		}
		return order_ports_out.GetOrderResult{}, fmt.Errorf(
			"scan errros: %w", err,
		)
	}

	if len(rawItems) != 0 {
		if err := json.Unmarshal(rawItems, &orderModel.Items); err != nil {
			return order_ports_out.GetOrderResult{}, fmt.Errorf(
				"json unmarshal: %w", err,
			)
		}
	}

	order := orderModel.ToDomain()

	return order_ports_out.NewGetOrderResult(order), nil
}
