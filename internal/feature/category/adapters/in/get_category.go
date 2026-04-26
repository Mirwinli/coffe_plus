package category_adapters_in

import (
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	categories_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

type GetCategoryResponse CategoryDTOResponse

// GetCategory godoc
// @Summary Отримання категорії
// @Desciption Отримання категорії по ID
// @Tags category
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Produce json
// @Success 200 {object} GetCategoryResponse "Катугорія"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /category/{id} [get]
func (h *CategoryHTTPHandler) GetCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get category id",
		)
		return
	}

	in := categories_ports_in.NewGetCategoryParams(id)
	category, err := h.categoryService.GetCategory(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get category",
		)
		return
	}

	response := GetCategoryResponse(categoryDTOFromDomain(category.Category))

	responseHandler.JSONResponse(response, http.StatusOK)
}
