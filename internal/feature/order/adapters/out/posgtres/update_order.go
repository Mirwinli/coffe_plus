package order_adapters_out_posgtres

import (
	"context"
	"fmt"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
)

func (r *OrderRepository) UpdateOrder(
	ctx context.Context,
	in order_ports_out.UpdateOrderParams,
) (order_ports_out.UpdateOrderResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `UPDATE coffe_plus.orders
			  SET
			     status = $1,
				 version = version + 1
			  WHERE id = $2 AND version = $3;
				 `

	cmd, err := r.pool.Exec(ctx, query, in.Order.Status, in.Order.ID, in.Order.Version)
	if err != nil {
		return order_ports_out.UpdateOrderResult{}, fmt.Errorf(
			"exec error: %w", err,
		)
	}

	if cmd.RowsAffected() == 0 {
		return order_ports_out.UpdateOrderResult{}, fmt.Errorf(
			"order concurrently accessed: %w",
			core_errors.ErrConflict,
		)
	}

	paramsOrder := order_ports_out.NewGetOrderParams(in.Order.ID)
	order, err := r.GetOrder(ctx, paramsOrder)
	if err != nil {
		return order_ports_out.UpdateOrderResult{}, fmt.Errorf(
			"get order from repository: %w", err,
		)
	}

	return order_ports_out.NewUpdateOrderResult(order.Order), nil
}
