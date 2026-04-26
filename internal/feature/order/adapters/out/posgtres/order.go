package order_adapters_out_posgtres

import (
	"context"
	"fmt"

	order_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/order/ports/out"
	"github.com/jackc/pgx/v5"
)

func (r *OrderRepository) SaveOrder(
	ctx context.Context,
	in order_ports_out.SaveOrderParams,
) (order_ports_out.SaveOrderResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return order_ports_out.SaveOrderResult{}, fmt.Errorf(
			"begin transaction: %w",
			err,
		)
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO coffe_plus.orders (id,customer_id, status, total_price, created_at) 
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id,version,customer_id, status, total_price, created_at; 
			  `

	var orderModel OrderModel
	var version int64
	if err := tx.QueryRow(
		ctx,
		query,
		in.Order.ID,
		in.Order.OrderReceiver.CustomerID,
		in.Order.Status,
		in.Order.Price,
		in.Order.CreatedAt,
	).Scan(
		&orderModel.ID,
		&version,
		&orderModel.OrderReceiver.CustomerID,
		&orderModel.Status,
		&orderModel.Price,
		&orderModel.CreatedAt,
	); err != nil {
		return order_ports_out.SaveOrderResult{}, fmt.Errorf(
			"insert order: %w",
			err,
		)
	}

	orderModel.Version = int(version)

	batch := &pgx.Batch{}

	itemQuery := `INSERT INTO coffe_plus.order_items (id,order_id,product_id,product_name,quantity,price_at_time,image_url)
				  VALUES ($1, $2, $3, $4, $5, $6,$7)
				  RETURNING id, order_id, product_id, product_name, quantity, price_at_time,image_url; 
				  `

	for _, item := range in.Order.Items {
		batch.Queue(itemQuery, item.ID, orderModel.ID, item.ProductID, item.Name, item.Quantity, item.Price_Per_Unit, item.ImageURL)
	}

	batchResult := tx.SendBatch(ctx, batch)

	var itemsModel []OrderItemModel
	for range in.Order.Items {
		rows, err := batchResult.Query()
		if err != nil {
			return order_ports_out.SaveOrderResult{}, fmt.Errorf(
				"query batch result: %w",
				err,
			)
		}

		for rows.Next() {
			var itemModel OrderItemModel
			if err := rows.Scan(
				&itemModel.ID,
				&itemModel.OrderID,
				&itemModel.ProductID,
				&itemModel.Name,
				&itemModel.Quantity,
				&itemModel.Price_Per_Unit,
				&itemModel.ImageURL,
			); err != nil {
				return order_ports_out.SaveOrderResult{}, fmt.Errorf(
					"scan batch result: %w",
					err,
				)
			}
			itemsModel = append(itemsModel, itemModel)
		}
		rows.Close()
	}

	if err := batchResult.Close(); err != nil {
		return order_ports_out.SaveOrderResult{}, fmt.Errorf(
			"close batch: %w",
			err,
		)
	}

	if err := tx.Commit(ctx); err != nil {
		return order_ports_out.SaveOrderResult{}, fmt.Errorf(
			"commit transaction: %w",
			err,
		)
	}

	orderModel.Items = itemsModel

	domainOrder := orderModel.ToDomain()
	domainOrder.OrderReceiver = in.Order.OrderReceiver
	return order_ports_out.NewSaveOrderResult(domainOrder), nil
}
