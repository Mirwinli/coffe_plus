package categories_ports_out

import "github.com/Mirwinli/coffe_plus/internal/core/domain"

type GetAllCategoriesParams struct {
	Limit  *int
	Offset *int
}

func NewGetAllCategoriesParams(limit, offset *int) GetAllCategoriesParams {
	return GetAllCategoriesParams{
		Limit:  limit,
		Offset: offset,
	}
}

type GetAllCategoriesResult struct {
	Categories []domain.Category
}

func NewGetAllCategoriesResult(c []domain.Category) GetAllCategoriesResult {
	return GetAllCategoriesResult{
		Categories: c,
	}
}
