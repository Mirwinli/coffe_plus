package adapters_in_auth_transport_http

import (
	"time"

	"github.com/google/uuid"
)

type UserDTOResponse struct {
	UserID      uuid.UUID `json:"user_id"      example:"1dsa-123s-1s"`
	Version     int       `json:"version"      example:"1"`
	FirstName   string    `json:"first_name"   example:"First Name"`
	LastName    string    `json:"last_name"    example:"Last Name"`
	PhoneNumber *string   `json:"phone_number" example:"+380974526180"`
	Email       string    `json:"email"        example:"email@gmail.com"`
	Role        string    `json:"role"         example:"common"`
	CreatedAt   time.Time `json:"created_at"   example:"0001-01-01T00:00:00Z"`
	AccessToken string    `json:"access_token" example:"access_token"`
}
