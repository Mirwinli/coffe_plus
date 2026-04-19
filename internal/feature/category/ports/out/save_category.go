package categories_ports_out

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
)

type SaveCategoryParams struct {
	Category domain.Category
}

func NewSaveCategoryParams(c domain.Category) SaveCategoryParams {
	return SaveCategoryParams{
		Category: c,
	}
}

type SaveCategoryResult struct {
	Category domain.Category
}

func NewSaveCategoryResult(category domain.Category) SaveCategoryResult {
	return SaveCategoryResult{
		Category: category,
	}
}
