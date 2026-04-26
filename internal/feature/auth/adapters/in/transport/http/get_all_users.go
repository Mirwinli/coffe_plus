package adapters_in_auth_transport_http

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
)

type GetAllUsersResponse []UserDTOResponse

// GetAllUsers godoc
// @Summary Всі користувачі
// @Description Вивести всіх користувачів
// @Description Only for admins
// @Tags users
// @Security BearerAuth
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} GetAllUsersResponse "Користувачі"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /admin/users [get]
func (h *AuthHTTPHandler) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandeler := core_http_response.NewHTTPResponseHandler(log, rw)

	offset, limit, err := core_http_request.GetOffsetLimitQueryParams(r)
	if err != nil {
		responseHandeler.ErrorResponse(
			err,
			"failed to get limit/offset query params",
		)
		return
	}

	in := auth_ports_in.NewGetAllUsersParams(limit, offset)
	getAllUsersResult, err := h.authService.GetAllUsers(ctx, in)
	if err != nil {
		responseHandeler.ErrorResponse(
			err,
			"failed to get all users",
		)
		return
	}

	response := GetAllUsersResponse(DTOsFromDomains(getAllUsersResult.Users))

	responseHandeler.JSONResponse(response, http.StatusOK)
}
