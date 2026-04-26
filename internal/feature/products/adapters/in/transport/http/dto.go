package products_adapters_in_products_transport_http

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type ProductDTOResponse struct {
	ID          uuid.UUID    `json:"id" example:"ba930185-467f-4031-b1bd-abf4899dffer"`
	Version     int          `json:"version" example:"1"`
	Name        string       `json:"name" example:"Pizza"`
	Description *string      `json:"description" example:"this pizza to hot"`
	Price       domain.Money `json:"price" example:"123.0"`
	IsAvailable bool         `json:"is_available" example:"true"`
	CategoryID  uuid.UUID    `json:"category_id" example:"ba930185-467f-4031-b1bd-abf4899dffer"`
	ImageURL    string       `json:"image_url" example:"https:cloudinary/image13"`
	PublicID    string       `json:"public_id" exampel:"ba930185-467f-4031-b1bd-abf4899dffer"`
}

func productDTOsFromDomains(products []domain.Product) []ProductDTOResponse {
	dtos := make([]ProductDTOResponse, len(products))
	for i, product := range products {
		dtos[i] = productDTOFromDomain(product)
	}

	return dtos
}

func productDTOFromDomain(product domain.Product) ProductDTOResponse {
	return ProductDTOResponse{
		ID:          product.ID,
		Version:     product.Version,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsAvailable: product.IsAvailable,
		CategoryID:  product.CategoryID,
		ImageURL:    product.ImageURL,
		PublicID:    product.PublicID,
	}
}
