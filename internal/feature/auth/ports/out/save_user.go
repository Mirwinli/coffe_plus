package auth_ports_out

import (
	"time"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type SaveUserAuthParams struct {
	User domain.User
}

func NewSaveUserAuthParams(user domain.User) SaveUserAuthParams {
	return SaveUserAuthParams{
		User: user,
	}
}

type SaveUserAuthResult struct {
	ID          uuid.UUID
	Version     int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Role        string
	CreatedAt   time.Time
}

func NewSaveUserAuthResult(
	id uuid.UUID,
	version int,
	firstName string,
	lastName string,
	email string,
	phoneNumber string,
	role string,
	time time.Time,
) SaveUserAuthResult {
	return SaveUserAuthResult{
		ID:          id,
		Version:     version,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        role,
		CreatedAt:   time,
	}
}
