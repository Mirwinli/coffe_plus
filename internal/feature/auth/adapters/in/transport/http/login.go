package adapters_in_auth_transport_http

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
)

type AuthLoginRequest struct {
	Email    string `json:"email"    validate:"required,min=3,max=100" example:"email@gmail.com"`
	Password string `json:"password" validate:"required,min=6,max=100" example:"123password"`
}

type AuthLoginResponse UserDTOResponse

// Login godoc
// @Summary     Вхід
// @Description Вхід вже зареєстрованого користувача в систему
// @Tags        auth
// @Accept 		json
// @Produce     json
// @Param 		request body AuthLoginRequest true "Login тіло запиту"
// @Success     200 {object} AuthLoginResponse "Користувач"
// @Header      200 {string} Set-Cookie "refresh_token=...; Path=/; SameSite=Strict; Max-Age=2592000; HttpOnly; Secure"
// @Failure     400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure     401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure     404 {object} core_http_response.ErrorResponse "Not found"
// @Failure     500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router      /auth/login [post]
func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request AuthLoginRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	ip := getIP(r)
	userAgent := r.UserAgent()

	in := auth_ports_in.NewLoginAuthParams(
		request.Email,
		request.Password,
		userAgent,
		&ip,
	)

	result, err := h.authService.Login(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to login",
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

	response := AuthLoginResponse{
		UserID:      result.ID,
		Version:     result.Version,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Role:        result.Role,
		AccessToken: result.AccessToken,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
