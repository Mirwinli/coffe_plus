package adapters_in_auth_transport_http

import (
	"net/http"

	core_contextKeys "github.com/Mirwinli/coffe_plus/internal/core/contextKeys"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	"github.com/google/uuid"
)

type GetMyselfUserResponse UserDTOResponse

// GetMyselfUser godoc
// @Summary Отримати себе
// @Description Отримати свого користувача
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} GetMyselfUserResponse "Користувач"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /profile [get]
func (h *AuthHTTPHandler) GetMyselfUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandeler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := ctx.Value(core_contextKeys.UserIDCtxKey).(uuid.UUID)
	if !ok {
		responseHandeler.ErrorResponse(
			core_errors.ErrInvalidArgument,
			"`userID` required",
		)
		return
	}

	in := auth_ports_in.NewGetUserParams(userID)
	getMyselfUserResult, err := h.authService.GetUser(ctx, in)
	if err != nil {
		responseHandeler.ErrorResponse(
			err,
			"failed to get myself user",
		)
	}

	response := GetMyselfUserResponse(DTOFromDomain(getMyselfUserResult.User))

	responseHandeler.JSONResponse(response, http.StatusOK)
}
