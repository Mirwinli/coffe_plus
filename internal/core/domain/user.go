package domain

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

const (
	versionUnitialized = -1
)

type User struct {
	ID          uuid.UUID
	Version     int
	Password    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber *string
	Role        string
	CreatedAt   time.Time
}

func NewUserUninitialized(
	firstName string,
	lastName string,
	password string,
	email string,
	phoneNumber *string,
) User {
	return User{
		ID:          uuid.New(),
		Version:     versionUnitialized,
		Password:    password,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        "common",
		CreatedAt:   time.Now(),
	}
}

func NewUser(
	id uuid.UUID,
	version int,
	password string,
	firstName string,
	lastName string,
	email string,
	phoneNumber *string,
	role string,
	createdAt time.Time,
) User {
	return User{
		ID:          id,
		Version:     version,
		Password:    password,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        role,
		CreatedAt:   createdAt,
	}
}

func (u *User) Validate() error {
	firstNameLen := len([]rune(u.FirstName))
	if firstNameLen < 3 || firstNameLen > 100 {
		return fmt.Errorf(
			"first name must be 3 between 100 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	lastNameLen := len([]rune(u.LastName))
	if lastNameLen < 3 || lastNameLen > 100 {
		return fmt.Errorf(
			"last name must be 3 between 100 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if err := ValidateEmail(u.Email); err != nil {
		return fmt.Errorf("validate email: %w", err)
	}

	passwordLen := len([]rune(u.Password))
	if passwordLen < 6 || passwordLen > 100 {
		return fmt.Errorf(
			"password must be 6 between 100 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLen := len(*u.PhoneNumber)
		if phoneNumberLen < 10 || phoneNumberLen > 13 {
			return fmt.Errorf(
				"phone number must be between 10 and 13 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		if !phoneNumberRegex.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid phone number: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil

}

func ValidateEmail(email string) error {
	emailLen := len(email)

	if emailLen < 3 || emailLen > 254 {
		return fmt.Errorf(
			"email address must be 3 between 254 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if !emailRegex.MatchString(email) {
		return fmt.Errorf(
			"invalid email address: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
