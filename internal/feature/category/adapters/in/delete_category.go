package category_adapters_in

import (
	"errors"
	"net/http"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	categories_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

// DeleteCategory godoc
// @Summary Видалення категорії
// @Description Видалення категорії продуктів зі системи
// @Description Щоб видалити категорію потрібно видалити спершу всі продукти цієї категорії
// @Description Only for admins
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Tags category
// @Success 204
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 404 {object} core_http_response.ErrorResponse "Not found"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /cart/{id} [delete]
func (h *CategoryHTTPHandler) DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get category id path",
		)
		return
	}

	in := categories_ports_in.NewDeleteCategoryParams(id)
	if err = h.categoryService.DeleteCategory(ctx, in); err != nil {
		if errors.Is(err, core_errors.ErrForeignKeyViolation) {
			responseHandler.ErrorResponse(
				err,
				"products in category,please first delete products",
			)
			return
		}
		responseHandler.ErrorResponse(
			err,
			"failed to delete category",
		)
		return
	}

	responseHandler.NoContentResponse()
}
