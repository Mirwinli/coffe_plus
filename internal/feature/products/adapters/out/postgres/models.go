package products_adapters_out_postgres

import "github.com/google/uuid"

type ProductModel struct {
	ID          uuid.UUID
	Version     int
	Name        string
	Description *string
	Price       float64
	IsAvaible   bool
	CategoryID  uuid.UUID
	ImageURL    string
	PublicID    string
}
