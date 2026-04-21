package cart_ports_in

import "github.com/google/uuid"

type DeleteCartParams struct {
	ID uuid.UUID
}

func NewDeleteCartParams(id uuid.UUID) DeleteCartParams {
	return DeleteCartParams{
		ID: id,
	}
}
