package auth_postgres

import (
	"time"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type UserModel struct {
	ID           uuid.UUID
	Version      int
	PasswordHash string
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	Role         string
	CreatedAt    time.Time
}

type RefreshModel struct {
	UserID    uuid.UUID
	ExpiresAt time.Time
}

func domainFromModel(userModel UserModel) domain.User {
	return domain.User{
		ID:          userModel.ID,
		Version:     userModel.Version,
		Password:    userModel.PasswordHash,
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
		Role:        userModel.Role,
		CreatedAt:   userModel.CreatedAt,
	}
}
