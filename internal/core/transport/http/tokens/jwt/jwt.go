package core_http_jwt

import (
	"fmt"
	"time"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"UserID"`
	Role   string    `json:"Role"`
	jwt.RegisteredClaims
}

func NewClaims(userID uuid.UUID, role string, config Config) Claims {
	return Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
}

func CreateToken(userID uuid.UUID, role string, config Config) (string, error) {
	claims := NewClaims(userID, role, config)

	tokenUnSigned := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenUnSigned.SignedString([]byte(config.Secret))
	if err != nil {
		return "", fmt.Errorf(
			"signing JWT token: %v: %w",
			err,
			core_errors.ErrUnauthorized,
		)
	}

	return token, nil
}

func ParseToken(tokenString string, config Config) (*Claims, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				"unexpected signing method: %w",
				core_errors.ErrUnauthorized,
			)
		}
		return []byte(config.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf(
			"parsing JWT token: %v: %w",
			err,
			core_errors.ErrUnauthorized,
		)
	}

	if !token.Valid {
		return nil, fmt.Errorf(
			"invalid JWT token: %w",
			core_errors.ErrUnauthorized,
		)
	}
	return &claims, nil
}
