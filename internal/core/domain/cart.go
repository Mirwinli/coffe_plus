package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `json:"id" examle:"ba930185-467f-4031-b1bd-abf4899dffde"`
	Items     []Item    `json:"items"`
	Price     Money     `json:"total_price" example:"100.12"`
	CreatedAt time.Time `json:"created_at"  example:"2026-04-23T16:25:43.655439Z"`
	UpdatedAt time.Time `json:"updated_at"  example:"2026-05-23T16:25:43.655439Z"`
}

func NewCart(
	id uuid.UUID,
	items []Item,
	price Money,
	createdAt time.Time,
	updatedAt time.Time,
) Cart {
	return Cart{
		ID:        id,
		Items:     items,
		Price:     price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type Item struct {
	ProductID      uuid.UUID `json:"product_id" example:"ba930185-467f-4031-b1bd-abf4899dffec"`
	Name           string    `json:"name"   	example:"Pizza"`
	ImageURL       string    `json:"image_url"  example:"https:/cloudinary/image12"`
	Quantity       int       `json:"quantity"	example:"12"`
	Price_Per_Unit Money     `json:"price_per_unit" example:"12.10"`
}

func NewItem(
	productID uuid.UUID,
	quantity int,
	price_per_unit Money,
	name string,
	imageURL string,
) Item {
	return Item{
		ProductID:      productID,
		Quantity:       quantity,
		Price_Per_Unit: price_per_unit,
		Name:           name,
		ImageURL:       imageURL,
	}
}
