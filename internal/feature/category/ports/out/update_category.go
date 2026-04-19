package categories_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type UpdateCategoryParams struct {
	Category domain.Category
}

func NewUpdateCategoryParams(category domain.Category) UpdateCategoryParams {
	return UpdateCategoryParams{
		Category: category,
	}
}

type UpdateCategoryResult struct {
	Category domain.Category
}

func NewUpdateCategoryResult(category domain.Category) UpdateCategoryResult {
	return UpdateCategoryResult{
		Category: category,
	}
}
