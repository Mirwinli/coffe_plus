package categories_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
)

type CreateCategoryParams struct {
	Category domain.Category
}

func NewCreateCategoryParams(c domain.Category) CreateCategoryParams {
	return CreateCategoryParams{
		Category: c,
	}
}

type CreateCategoryResult struct {
	Category domain.Category
}

func NewCreateCategoryResult(category domain.Category) CreateCategoryResult {
	return CreateCategoryResult{
		Category: category,
	}
}
