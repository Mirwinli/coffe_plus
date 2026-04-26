package adapters_in_auth_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	core_http_types "github.com/Mirwinli/coffe_plus/internal/core/transport/http/types"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	"github.com/google/uuid"
)

type PatchUserSwaggerRequest struct {
	FirstName   *string `json:"first_name"   example:"Max"`
	LastName    *string `json:"last_name"    example:"Trump"`
	Password    *string `json:"password"     example:"123456hs"`
	PhoneNumber *string `json:"phone_number" example:"+380974526180"`
}

type PatchUserRequest struct {
	FirstName   core_http_types.Nullable[string] `json:"first_name"  swaggertype:"string" example:"Max"`
	LastName    core_http_types.Nullable[string] `json:"last_name"	 swaggertype:"string" example:"Trump"`
	Password    core_http_types.Nullable[string] `json:"password"	 swaggertype:"string" example:"123456hs"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" swaggertype:"string" example:"+380974526180"`
}

type PatchUserResponse struct {
	UserID      uuid.UUID `json:"user_id"      example:"1dsa-123s-1s"`
	Version     int       `json:"version"      example:"1"`
	FirstName   string    `json:"first_name"   example:"First Name"`
	LastName    string    `json:"last_name"    example:"Last Name"`
	PhoneNumber string    `json:"phone_number" example:"+380974526180"`
	Email       string    `json:"email"        example:"email@gmail.com"`
	Role        string    `json:"role"         example:"common"`
	CreatedAt   time.Time `json:"created_at"   example:"0001-01-01T00:00:00Z"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FirstName.Set {
		if r.FirstName.Value == nil {
			return fmt.Errorf(
				"first name must not be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}
		firstNameLen := len([]rune(*r.FirstName.Value))
		if firstNameLen < 1 || firstNameLen > 100 {
			return fmt.Errorf(
				"first name must be between 1 and 100 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.LastName.Set {
		if r.LastName.Value == nil {
			return fmt.Errorf(
				"last name must not be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}
		lastNameLen := len([]rune(*r.LastName.Value))
		if lastNameLen < 1 || lastNameLen > 100 {
			return fmt.Errorf(
				"last name must be between 1 and 100 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.Password.Set {
		if r.Password.Value == nil {
			return fmt.Errorf(
				"password must not be null: %w",
			)
		}
		passwordLen := len([]rune(*r.Password.Value))
		if passwordLen < 6 || passwordLen > 100 {
			return fmt.Errorf(
				"password must be between 6 and 100 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value == nil {
			return fmt.Errorf(
				"phone number must not be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}
		phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
		if phoneNumberLen < 10 || phoneNumberLen > 13 {
			return fmt.Errorf(
				"phonen number must be between 10 and 13 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

// PatchUser godoc
// @Summary Зміна користувача
// @Description Зміна данних користувача
// @Description Номеру телефону,Прізвище або Імя,Пароль
// @Description Правила зміни полів:
// @Description 1. **Поле не передано**: `description` ігноруєтся, значення в БД не змінюєтся
// @Description 2. **Явно передано значення**: `description`: "Вийти погуляти в 6:30 з собакою" - змінюєьтся поле в БД
// @Description Обмеження: Ніякі поля не можуть бути null
// @Security BearerAuth
// @Tags auth
// @Accept json
// @Produce json
// @Param request body PatchUserSwaggerRequest true "PatchUser тіло запиту"
// @Success 200 {object} PatchUserResponse "Змінений користувач"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 404 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /auth/user [patch]
func (h *AuthHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	userID, ok := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)
	if !ok {
		responseHandler.ErrorResponse(
			core_errors.ErrUnauthorized,
			"failed to get user id",
		)
		return
	}

	in := auth_ports_in.NewPatchUserParams(domain.NewPatchUser(
		request.FirstName.ToDomain(),
		request.LastName.ToDomain(),
		request.Password.ToDomain(),
		request.PhoneNumber.ToDomain(),
	),
		userID,
	)

	patchedUser, err := h.authService.PatchUser(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
		return
	}

	user := patchedUser.User
	response := PatchUserResponse{
		UserID:      user.ID,
		Version:     user.Version,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
