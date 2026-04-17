package auth_ports_out

import (
	"time"

	"github.com/google/uuid"
)

type SaveRefreshTokenAuthParams struct {
	UserID     uuid.UUID
	TokenHash  string
	DeviceName string
	IpAddress  *string
	CreatedAt  time.Time
	ExpiresAt  time.Time
}

func NewSaveRefreshTokenAuthParams(
	refreshToken string,
	userID uuid.UUID,
	deviceName string,
	ipAddress *string,
	expiresAt time.Time,
) SaveRefreshTokenAuthParams {
	return SaveRefreshTokenAuthParams{
		TokenHash:  refreshToken,
		UserID:     userID,
		DeviceName: deviceName,
		IpAddress:  ipAddress,
		CreatedAt:  time.Now(),
		ExpiresAt:  expiresAt,
	}

}
