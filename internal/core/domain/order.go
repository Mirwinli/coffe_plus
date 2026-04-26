package domain

import (
	"fmt"
	"time"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID     `json:"id" example:"3d1c7344-a98d-4a6a-847a-caa29238cfer"`
	Version       int           `json:"version" example:"1"`
	Items         []OrderItem   `json:"items"`
	Status        string        `json:"status" example:"cooking"`
	Price         Money         `json:"total_price" example:"120.10"`
	OrderReceiver OrderReceiver `json:"order_receiver"`
	CreatedAt     time.Time     `json:"created_at" example:"2026-04-23T16:25:43.655439Z"`
}

func NewOrderUnitialized(
	items []OrderItem,
	price Money,
	orderReceiver OrderReceiver,
) Order {
	return Order{
		ID:            uuid.New(),
		Version:       versionUnitialized,
		Items:         items,
		Status:        StatusCreated,
		Price:         price,
		OrderReceiver: orderReceiver,
		CreatedAt:     time.Now(),
	}
}

func NewOrder(
	id uuid.UUID,
	version int,
	items []OrderItem,
	price Money,
	status string,
	orderReceiver OrderReceiver,
	createdAt time.Time,
) Order {
	return Order{
		ID:            id,
		Version:       version,
		Items:         items,
		Status:        status,
		Price:         price,
		OrderReceiver: orderReceiver,
		CreatedAt:     createdAt,
	}
}

func (o Order) ValidateStatusTransitons(futureStatus string) error {
	allowed, ok := AllowedTransitions[o.Status]
	if !ok {
		return fmt.Errorf(
			"unknow status: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	for _, v := range allowed {
		if v == futureStatus {
			return nil
		}
	}

	return fmt.Errorf(
		"transition from %s to %s is not allowed: %w",
		o.Status, futureStatus, core_errors.ErrInvalidArgument,
	)
}

type OrderReceiver struct {
	CustomerID  uuid.UUID `json:"customer_id" example:"3d1c7344-a98d-4a6a-847a-caa29238cfer"`
	PhoneNumber string    `json:"phone_number" example:"+380974526180"`
	Email       string    `json:"email"        example:"email@gmail.com"`
	FirstName   string    `json:"first_name"   example:"Max"`
	LastName    string    `json:"last_name"	   example:"Trump"`
}

func NewOrderReceiver(
	customerID uuid.UUID,
	phoneNumber string,
	email string,
	firstName string,
	lastName string,
) OrderReceiver {
	return OrderReceiver{
		CustomerID:  customerID,
		PhoneNumber: phoneNumber,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
	}
}

type OrderItem struct {
	ID             uuid.UUID `json:"id" example:"3d1c7344-a98d-4a6a-847a-caa29238cfes"`
	ProductID      uuid.UUID `json:"product_id" example:"3d1c7344-a98d-4a6a-847a-caa29238cfer"`
	Name           string    `json:"name" example:"pizza"`
	ImageURL       string    `json:"image_url" example:"https:cloudinry/image12"`
	Quantity       int       `json:"quantity"  example:"12"`
	Price_Per_Unit Money     `json:"price_at_time" example:"21"`
}

func NewOrderItemUnitialized(
	productID uuid.UUID,
	name string,
	imageURL string,
	quantity int,
	price_Per_Unit Money,
) OrderItem {
	return OrderItem{
		ID:             uuid.New(),
		ProductID:      productID,
		Name:           name,
		ImageURL:       imageURL,
		Quantity:       quantity,
		Price_Per_Unit: price_Per_Unit,
	}
}

func NewOrderItem(
	id uuid.UUID,
	productID uuid.UUID,
	name string,
	imageURL string,
	quantity int,
	price_Per_Unit Money,
) OrderItem {
	return OrderItem{
		ID:             id,
		ProductID:      productID,
		Name:           name,
		ImageURL:       imageURL,
		Quantity:       quantity,
		Price_Per_Unit: price_Per_Unit,
	}
}

func NewOrderItems(items []Item) []OrderItem {
	orderItems := make([]OrderItem, len(items))

	for i, item := range items {
		orderItems[i] = NewOrderItemUnitialized(
			item.ProductID,
			item.Name,
			item.ImageURL,
			item.Quantity,
			item.Price_Per_Unit,
		)
	}

	return orderItems
}

func (o *Order) Validate() error {
	if o.Items == nil {
		return fmt.Errorf(
			"items cannot be nil: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if o.Status == "" {
		return fmt.Errorf(
			"status cannot be empty: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if o.Price.IsZero() {
		return fmt.Errorf(
			"price cannot be zero: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if o.OrderReceiver.CustomerID == uuid.Nil {
		return fmt.Errorf(
			"order_receiver.CustomerID cannot be nil: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if o.OrderReceiver.PhoneNumber == "" {
		return fmt.Errorf(
			"order_receiver.PhoneNumber cannot be empty: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if o.OrderReceiver.Email == "" {
		return fmt.Errorf(
			"order_receiver.Email cannot be empty: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if o.OrderReceiver.FirstName == "" {
		return fmt.Errorf(
			"order_receiver.FirstName cannot be empty: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if o.OrderReceiver.LastName == "" {
		return fmt.Errorf(
			"order_receiver.LastName cannot be empty: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
