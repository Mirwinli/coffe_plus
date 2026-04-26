package auth_ports_in

import (
	"fmt"
	"time"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
)

type LoginAuthParams struct {
	Email      string
	Password   string
	DeviceName string
	IpAddress  *string
}

func (p *LoginAuthParams) Validate() error {
	emailLen := len(p.Email)

	if emailLen < 3 || emailLen > 254 {
		return fmt.Errorf(
			"email address must be 3 between 254 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if !emailRegex.MatchString(p.Email) {
		return fmt.Errorf(
			"invalid email address: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	passwordLen := len([]rune(p.Password))
	if passwordLen < 6 || passwordLen > 100 {
		return fmt.Errorf(
			"password must be 6 between 100 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func NewLoginAuthParams(
	email string,
	password string,
	DeviceName string,
	IpAddress *string,
) LoginAuthParams {
	return LoginAuthParams{
		Email:      email,
		Password:   password,
		DeviceName: DeviceName,
		IpAddress:  IpAddress,
	}
}

type LoginAuthResult struct {
	ID           uuid.UUID
	Version      int
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	Role         string
	CreatedAt    time.Time
	AccessToken  string
	RefreshToken string
}

func NewLoginAuthResult(
	id uuid.UUID,
	version int,
	firstName string,
	lastName string,
	email string,
	phoneNumber string,
	role string,
	createdAt time.Time,
	accessToken string,
	refreshToken string,
) LoginAuthResult {
	return LoginAuthResult{
		ID:           id,
		Version:      version,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PhoneNumber:  phoneNumber,
		Role:         role,
		CreatedAt:    createdAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
