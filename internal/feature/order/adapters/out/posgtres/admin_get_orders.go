package order_adapters_out_posgtres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (r *OrderRepository) AdminGetOrders(
	ctx context.Context,
	in order_ports_out.AdminGetOrdersParams,
) (order_ports_out.AdminGetOrdersResult, error) {
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
			 %s
			 GROUP BY o.id, u.id, u.phone_number, u.email, u.first_name, u.last_name
			 %s`

	orderBy := "ORDER BY o.created_at DESC"
	where := ""

	args := []any{}
	if in.Status != nil {
		where = "WHERE o.status = $1"
		args = append(args, *in.Status)
		if *in.Status == domain.StatusCreated {
			orderBy = "ORDER BY o.created_at ASC"
		}
	}

	finalyQuery := fmt.Sprintf(
		query,
		where,
		orderBy,
	)

	rows, err := r.pool.Query(ctx, finalyQuery, args...)
	if err != nil {
		return order_ports_out.AdminGetOrdersResult{}, fmt.Errorf(
			"select error: %w", err,
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
			return order_ports_out.AdminGetOrdersResult{}, fmt.Errorf(
				"scan error: %w", err,
			)
		}
		if len(itemsRaw) != 0 {
			if err := json.Unmarshal(itemsRaw, &orderModel.Items); err != nil {
				return order_ports_out.AdminGetOrdersResult{}, fmt.Errorf(
					"unmarshal items: %w", err,
				)
			}
		}
		orders = append(orders, orderModel)
	}

	domains := orderModelsToDomains(orders)

	return order_ports_out.NewAdminGetOrdersResult(domains), nil
}
