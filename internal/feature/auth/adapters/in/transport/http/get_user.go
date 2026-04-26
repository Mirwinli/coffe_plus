package adapters_in_auth_transport_http

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
)

type GetUserResponse UserDTOResponse

// GetUser godoc
// @Summary Отримати користувача
// @Description Отримати користувача за його ID
// @Description Only for admins
// @Tags users
// @Security BearerAuth
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} GetUserResponse "Користувач"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /admin/users/{id} [get]
func (h *AuthHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `userID` path value",
		)
		return
	}

	in := auth_ports_in.NewGetUserParams(userID)
	user, err := h.authService.GetUser(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}

	response := GetUserResponse(DTOFromDomain(user.User))

	responseHandler.JSONResponse(response, http.StatusOK)
}
