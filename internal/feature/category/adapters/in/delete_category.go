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
