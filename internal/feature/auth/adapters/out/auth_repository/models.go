package auth_postgres

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID           uuid.UUID
	Version      int
	PasswordHash string
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  *string
	Role         string
	CreatedAt    time.Time
}

type RefreshModel struct {
	UserID    uuid.UUID
	ExpiresAt time.Time
}
