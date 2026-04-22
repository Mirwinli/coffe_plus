package auth_ports_in

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
)

var (
	emailRegex       = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	phoneNumberRegex = regexp.MustCompile(`^(\+?38)?0\d{9}$`)
)

type RegisterAuthParams struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
	DeviceName  string
	IpAddress   *string
}

func (p *RegisterAuthParams) Validate() error {
	firstNameLen := len([]rune(p.FirstName))
	if firstNameLen < 3 || firstNameLen > 100 {
		return fmt.Errorf(
			"first name must be 3 between 100 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	lastNameLen := len([]rune(p.LastName))
	if lastNameLen < 3 || lastNameLen > 100 {
		return fmt.Errorf(
			"last name must be 3 between 100 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

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

	phoneNumberLen := len(p.PhoneNumber)
	if phoneNumberLen < 10 || phoneNumberLen > 13 {
		return fmt.Errorf(
			"phone number must be between 10 and 13 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if !phoneNumberRegex.MatchString(p.PhoneNumber) {
		return fmt.Errorf(
			"invalid phone number: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func NewRegisterAuthParams(
	firstName string,
	lastName string,
	email string,
	phoneNumber string,
	password string,
	deviceName string,
	ipAddress *string,
) RegisterAuthParams {
	return RegisterAuthParams{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
		DeviceName:  deviceName,
		IpAddress:   ipAddress,
	}
}

type RegisterAuthResult struct {
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

func NewRegisterAuthResult(
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
) RegisterAuthResult {
	return RegisterAuthResult{
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
