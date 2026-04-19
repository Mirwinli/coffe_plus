package categories_ports_in

import "github.com/google/uuid"

type DeleteCategoryParams struct {
	ID uuid.UUID
}

func NewDeleteCategoryParams(id uuid.UUID) DeleteCategoryParams {
	return DeleteCategoryParams{
		ID: id,
	}
}
