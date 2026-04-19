package adapters_in_auth_transport_http

import (
	"errors"
	"net/http"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
)

type AuthRefreshResponse struct {
	AccessToken string `json:"access_token" example:"access_token"`
}

// Refresh godoc
// @Summary      Зміна access та refresh токенів
// @Description  Видалення та зміна refresh токен та створення нового access токен
// @Tags         auth
// @Produce      json
// @Param Cookie header string true "refresh_token=your_token_here"
// @Success      200 {object} AuthLoginResponse "Success" {header:string:Set-Cookie:"Session cookie"}
// @Failure      500 {object} core_http_response.ErrorResponse "Internal server error"
// @Failure      401 {object} core_http_response.ErrorResponse "Invalid token"
// @Router       /auth/refresh [post]
func (h *AuthHTTPHandler) Refresh(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	oldRefresh, err := r.Cookie("refresh_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			responseHandler.ErrorResponse(
				core_errors.ErrUnauthorized,
				"refresh_token cookie not found",
			)
			return
		}
		responseHandler.ErrorResponse(
			err,
			"failed to get refresh token",
		)
		return
	}

	ip := getIP(r)
	userAgent := r.UserAgent()

	in := auth_ports_in.NewRefreshAuthParams(oldRefresh.Value, userAgent, &ip)
	result, err := h.authService.Refresh(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to refresh",
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

	response := AuthRefreshResponse{
		AccessToken: result.AccessToken,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
