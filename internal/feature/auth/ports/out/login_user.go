package auth_ports_out

import "github.com/google/uuid"

type LoginUserAuthParams struct {
	Email    string
	Password string
}

func NewLoginUserAuthParams(
	email string,
	password string,
) LoginUserAuthParams {
	return LoginUserAuthParams{
		Email:    email,
		Password: password,
	}
}

type LoginUserAuthResult struct {
	ID           uuid.UUID
	Version      int
	PasswordHash string
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	Role         string
}

func NewLoginUserAuthResult(
	id uuid.UUID,
	version int,
	passwordHash string,
	firstName string,
	lastName string,
	email string,
	phoneNumber string,
	role string,
) LoginUserAuthResult {
	return LoginUserAuthResult{
		ID:           id,
		Version:      version,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Role:         role,
		PasswordHash: passwordHash,
	}
}
