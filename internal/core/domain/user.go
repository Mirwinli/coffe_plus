package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	PhoneNumber string
	Role        string
	CreatedAt   time.Time
}

func NewUserUninitialized(
	firstName string,
	lastName string,
	password string,
	email string,
	phoneNumber string,
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
	phoneNumber string,
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

	phoneNumberLen := len(u.PhoneNumber)
	if phoneNumberLen < 10 || phoneNumberLen > 13 {
		return fmt.Errorf(
			"phone number must be between 10 and 13 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if !phoneNumberRegex.MatchString(u.PhoneNumber) {
		return fmt.Errorf(
			"invalid phone number: %w",
			core_errors.ErrInvalidArgument,
		)
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

type PatchUser struct {
	FirstName   Nullable[string]
	LastName    Nullable[string]
	Password    Nullable[string]
	PhoneNumber Nullable[string]
}

func NewPatchUser(
	firstName Nullable[string],
	lastName Nullable[string],
	password Nullable[string],
	phoneNumber Nullable[string],
) PatchUser {
	return PatchUser{
		FirstName:   firstName,
		LastName:    lastName,
		Password:    password,
		PhoneNumber: phoneNumber,
	}
}

func (u *PatchUser) Validate() error {
	if u.FirstName.Set && u.FirstName.Value == nil {
		return fmt.Errorf(
			"first name cannot be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if u.LastName.Set && u.LastName.Value == nil {
		return fmt.Errorf(
			"last name cannot be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if u.Password.Set && u.Password.Value == nil {
		return fmt.Errorf(
			"password cannot be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if u.PhoneNumber.Set && u.PhoneNumber.Value == nil {
		return fmt.Errorf(
			"phone number cannot be NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch PatchUser) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate patch: %w", err)
	}

	tmp := *u

	if patch.FirstName.Set {
		tmp.FirstName = *patch.FirstName.Value
	}

	if patch.LastName.Set {
		tmp.LastName = *patch.LastName.Value
	}

	if patch.Password.Set {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(*patch.Password.Value), 12)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}

		tmp.Password = string(passwordHash)
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = *patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate tmp: %w", err)
	}

	*u = tmp

	return nil
}
