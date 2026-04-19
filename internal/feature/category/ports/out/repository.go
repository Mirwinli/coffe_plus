package categories_ports_out

import "context"

type CategoriesRepository interface {
	SaveCategory(
		ctx context.Context,
		in SaveCategoryParams,
	) (SaveCategoryResult, error)
	UpdateCategory(
		ctx context.Context,
		in UpdateCategoryParams,
	) (UpdateCategoryResult, error)
	DeleteCategory(
		ctx context.Context,
		in DeleteCategoryParams,
	) error
	GetCategory(
		ctx context.Context,
		in GetCategoryParams,
	) (GetCategoryResult, error)
	GetAllCategories(
		ctx context.Context,
		in GetAllCategoriesParams,
	) (GetAllCategoriesResult, error)
}
