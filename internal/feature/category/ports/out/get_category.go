package categories_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type GetCategoryParams struct {
	ID uuid.UUID
}

func NewGetCategoryParams(id uuid.UUID) GetCategoryParams {
	return GetCategoryParams{
		ID: id,
	}
}

type GetCategoryResult struct {
	Category domain.Category
}

func NewGetCategoryResult(c domain.Category) GetCategoryResult {
	return GetCategoryResult{
		Category: c,
	}
}
