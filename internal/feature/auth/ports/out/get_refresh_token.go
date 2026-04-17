package auth_ports_out

import (
	"time"

	"github.com/google/uuid"
)

type GetRefreshParams struct {
	RefreshToken string
	DeviceName   string
}

func NewGetRefreshParams(refreshToken string, deviceName string) GetRefreshParams {
	return GetRefreshParams{
		RefreshToken: refreshToken,
		DeviceName:   deviceName,
	}
}

type GetRefreshResult struct {
	UserID    uuid.UUID
	ExpiresAt time.Time
}

func NewGetRefreshResult(expiresAt time.Time, userID uuid.UUID) GetRefreshResult {
	return GetRefreshResult{
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
}
