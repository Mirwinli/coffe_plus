package auth_ports_out

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BlackListParams struct {
	IDAccess  uuid.UUID
	ExpiresAt *jwt.NumericDate
}

func NewBlackListParams(idAccess uuid.UUID, date *jwt.NumericDate) BlackListParams {
	return BlackListParams{
		IDAccess:  idAccess,
		ExpiresAt: date,
	}
}
