package category_adapters_in

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	categories_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

type GetAllCategoriesResponse []CategoryDTOResponse

func (h *CategoryHTTPHandler) GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	offset, limit, err := getOffsetLimitQueryParams(r)
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

func getOffsetLimitQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParam  = "limit"
		offsetQueryParam = "offset"
	)

	limit, err := core_http_request.GetIntQueryParams(r, limitQueryParam)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get `offset` query param: %w",
			err,
		)
	}

	offset, err := core_http_request.GetIntQueryParams(r, offsetQueryParam)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get `offset` query param: %w",
			err,
		)
	}

	return offset, limit, nil
}
