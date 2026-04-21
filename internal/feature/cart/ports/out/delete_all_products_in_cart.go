package cart_ports_out

import "github.com/google/uuid"

type DeleteCartParams struct {
	ID uuid.UUID
}

func NewDeleteCartParams(id uuid.UUID) DeleteCartParams {
	return DeleteCartParams{
		ID: id,
	}
}
