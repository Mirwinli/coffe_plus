package auth_ports_out

import (
	"github.com/golang-jwt/jwt/v5"
)

type BlackListParams struct {
	IDAccess  string
	ExpiresAt *jwt.NumericDate
}

func NewBlackListParams(idAccess string, date *jwt.NumericDate) BlackListParams {
	return BlackListParams{
		IDAccess:  idAccess,
		ExpiresAt: date,
	}
}
