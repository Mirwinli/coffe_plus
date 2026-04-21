package products_adapters_in_products_transport_http

import (
	"net/http"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	core_http_types "github.com/Mirwinli/coffe_plus/internal/core/transport/http/types"
	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	"github.com/google/uuid"
)

type PatchProductRequest struct {
	Name         core_http_types.Nullable[string]       `json:"name"`
	Description  core_http_types.Nullable[string]       `json:"description"`
	Price        core_http_types.Nullable[domain.Money] `json:"price"`
	CategoryID   core_http_types.Nullable[uuid.UUID]    `json:"category_id"`
	Is_Available core_http_types.Nullable[bool]         `json:"is_available"`
}

type PatchProductResponse ProductDTOResponse

func (h *ProductsHTTPHandler) PatchProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request PatchProductRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	id, err := core_http_request.GetUUIDPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get id from path",
		)
		return
	}

	productPatch := productPatchFromRequest(request)

	in := products_ports_in.NewPatchProductParams(id, productPatch)

	result, err := h.productsService.PatchProduct(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch product",
		)
		return
	}

	response := GetProductResponse(productDTOFromDomain(result.Product))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func productPatchFromRequest(request PatchProductRequest) domain.ProductPatch {
	return domain.NewProductPatch(
		request.Name.ToDomain(),
		request.Description.ToDomain(),
		request.Price.ToDomain(),
		request.CategoryID.ToDomain(),
		request.Is_Available.ToDomain(),
	)
}
