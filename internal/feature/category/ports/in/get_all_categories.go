package categories_ports_in

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type GetAllCategoriesParams struct {
	Offset *int
	Limit  *int
}

func NewGetAllCategoriesParams(offset *int, limit *int) GetAllCategoriesParams {
	return GetAllCategoriesParams{
		Offset: offset,
		Limit:  limit,
	}
}

type GetAllCategoriesResult struct {
	Categories []domain.Category
}

func NewGetAllCategoriesResult(categories []domain.Category) GetAllCategoriesResult {
	return GetAllCategoriesResult{
		Categories: categories,
	}
}
