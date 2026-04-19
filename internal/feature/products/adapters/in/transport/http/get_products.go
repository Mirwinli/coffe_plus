package products_adapters_in_products_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	"github.com/google/uuid"
)

type GetProductsResponse []ProductDTOResponse

func (h *ProductsHTTPHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, categoryID, err := getCategoryLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `limit` and `offset` query params",
		)
		return
	}

	in := products_ports_in.NewGetProductsParams(limit, offset, categoryID)
	result, err := h.productsService.GetProducts(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get products",
		)
		return
	}

	dtos := productDTOsFromDomains(result.Products)
	response := GetProductsResponse(dtos)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getCategoryLimitOffsetQueryParams(r *http.Request) (*int, *int, *uuid.UUID, error) {
	const (
		limitQueryParamKey    = "limit"
		categoryQueryParamKey = "category_id"
		offsetQueryParamKey   = "offset"
	)

	limit, err := core_http_request.GetIntQueryParams(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"get `limit` query params: %w",
			err,
		)
	}

	category, err := core_http_request.GetUUIDQueryParams(r, categoryQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"get `category` query params: %w",
		)
	}

	offset, err := core_http_request.GetIntQueryParams(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"get `offset` query params: %w",
			err,
		)
	}

	return limit, offset, category, nil
}
