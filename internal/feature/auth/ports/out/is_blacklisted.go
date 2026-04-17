package auth_ports_out

import "github.com/google/uuid"

type IsBlackListedParams struct {
	IDAccess uuid.UUID
}

func NewIsBlackListedParams(idAccess uuid.UUID) IsBlackListedParams {
	return IsBlackListedParams{
		IDAccess: idAccess,
	}

}
