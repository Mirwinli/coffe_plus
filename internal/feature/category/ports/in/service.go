package categories_ports_in

import "context"

type CategoryService interface {
	CreateCategory(
		ctx context.Context,
		in CreateCategoryParams,
	) (CreateCategoryResult, error)
	PatchCategory(
		ctx context.Context,
		in PatchCategoryParams,
	) (PatchCategoryResult, error)
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
