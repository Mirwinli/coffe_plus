package products_ports_out

import (
	"github.com/google/uuid"
)

type DeleteProductParams struct {
	ID uuid.UUID
}

func NewDeleteProductParams(id uuid.UUID) DeleteProductParams {
	return DeleteProductParams{
		ID: id,
	}
}
