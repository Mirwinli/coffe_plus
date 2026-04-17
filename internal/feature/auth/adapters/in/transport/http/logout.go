package auth_transport_http

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
)

// Logout godoc
// @Summary Вихід
// @Description Видалення refresh токену та скасування access токену у зареєстрованого користувача
// @Tags auth
// @Param Cookie header string true "refresh_token=your_token_here"
// @Header 200 {string} Set-Cookie "refresh_token=; Path=/;MaxAge=-1;HttpOnly;Secure"
// @Success 204 "Успішний вихід"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router /auth/logout [delete]
func (h *AuthHTTPHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get refresh token",
		)
		return
	}

	userAgent := r.UserAgent()

	in := auth_ports_in.NewLogoutAuthParams(userAgent, refreshToken.Value)

	if err = h.authService.Logout(ctx, in); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to logout",
		)
		return
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.Value,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(rw, cookie)

	responseHandler.NoContentResponse()
}
