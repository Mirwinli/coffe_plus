package category_adapters_in

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	categories_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

type GetAllCategoriesResponse []CategoryDTOResponse

// GetAllCategories
// @Summary Виведення всіх категорій
// @Description Виведення всіх категорій продуктів
// @Description Only for admins
// @Tags category
// @Security BearerAuth
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} GetAllCategoriesResponse "Всі категорії"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /category [get]
func (h *CategoryHTTPHandler) GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	offset, limit, err := core_http_request.GetOffsetLimitQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get offset and limit query params",
		)
		return
	}

	in := categories_ports_in.NewGetAllCategoriesParams(offset, limit)

	categories, err := h.categoryService.GetAllCategories(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get all categories",
		)
		return
	}

	response := GetAllCategoriesResponse(categoryDTOsFromDomains(categories.Categories))

	responseHandler.JSONResponse(response, http.StatusOK)
}
