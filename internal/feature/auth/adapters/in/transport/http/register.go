package adapters_in_auth_transport_http

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
)

type RegisterRequest struct {
	Email       string  `json:"email"       validate:"required,min=3,max=100"   example:"email@gmail.com"`
	Password    string  `json:"password"    validate:"required,min=6,max=100"   example:"123password"`
	FirstName   string  `json:"first_name"  validate:"required,min=3,max=100"   example:"First Name"`
	LastName    string  `json:"last_name"   validate:"required,min=3,max=100"   example:"Last Name"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=13" example:"+380974526180"`
}

type RegisterResponse UserDTOResponse

// Register godoc
// @Summary Реєстрація
// @Description Створення нового користувача в системі
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register тіло запиту"
// @Header 201 {string} Set-Cookie "refresh_token=...; Path=/; SameSite=Strict;MaxAge=;HttpOnly;Secure"
// @Success 201 {object} RegisterResponse "Користувач"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *AuthHTTPHandler) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request RegisterRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode request",
		)
		return
	}

	ip := getIP(r)
	userAgent := r.UserAgent()

	in := auth_ports_in.NewRegisterAuthParams(
		request.FirstName,
		request.LastName,
		request.Email,
		request.PhoneNumber,
		request.Password,
		userAgent,
		&ip,
	)

	result, err := h.authService.Register(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to register user",
		)
		return
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    result.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(h.JWTConfig.RefreshTokenTTL.Seconds()),
	}

	http.SetCookie(rw, cookie)

	response := RegisterResponse{
		UserID:      result.ID,
		Version:     result.Version,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		PhoneNumber: result.PhoneNumber,
		Email:       result.Email,
		Role:        result.Role,
		CreatedAt:   result.CreatedAt,
		AccessToken: result.AccessToken,
	}

	responseHandler.JSONResponse(response, http.StatusCreated)
}
