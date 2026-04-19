package category_adapters_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type CategoryDTOResponse struct {
	ID   uuid.UUID
	Name string
}

func categoryDTOFromDomain(category domain.Category) CategoryDTOResponse {
	return CategoryDTOResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func categoryDTOsFromDomains(categories []domain.Category) []CategoryDTOResponse {
	dtos := make([]CategoryDTOResponse, len(categories))

	for i, category := range categories {
		dtos[i] = categoryDTOFromDomain(category)
	}
	
	return dtos
}
