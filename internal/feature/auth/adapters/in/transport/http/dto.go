package adapters_in_auth_transport_http

import (
	"time"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type UserAndAccessDTOResponse struct {
	UserID      uuid.UUID `json:"user_id"      example:"1ba930185-467f-4031-b1bd-abf4899dffde"`
	Version     int       `json:"version"      example:"1"`
	FirstName   string    `json:"first_name"   example:"First Name"`
	LastName    string    `json:"last_name"    example:"Last Name"`
	PhoneNumber string    `json:"phone_number" example:"+380974526180"`
	Email       string    `json:"email"        example:"email@gmail.com"`
	Role        string    `json:"role"         example:"common"`
	CreatedAt   time.Time `json:"created_at"   example:"0001-01-01T00:00:00Z"`
	AccessToken string    `json:"access_token" example:"access_token"`
}

type UserDTOResponse struct {
	ID          uuid.UUID `json:"user_id"      example:"ba930185-467f-4031-b1bd-abf4899dffde"`
	Version     int       `json:"version"	   example:"1"`
	FirstName   string    `json:"first_name"   example:"Max"`
	LastName    string    `json:"last_name"	   example:"Trump"`
	Email       string    `json:"email"		   example:"email@gmail.com"`
	PhoneNumber string    `json:"phone_number" example:"+380974526180"`
	Role        string    `json:"role"		   example:"commom"`
	CreatedAt   time.Time `json:"created_at"   example:"0001-01-01T00:00:00Z"`
}

func DTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
	}
}

func DTOsFromDomains(users []domain.User) []UserDTOResponse {
	dtos := make([]UserDTOResponse, len(users))

	for i, user := range users {
		dtos[i] = DTOFromDomain(user)
	}

	return dtos
}
