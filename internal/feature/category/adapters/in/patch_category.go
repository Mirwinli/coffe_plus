package category_adapters_in

import (
	"fmt"
	"net/http"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	core_http_types "github.com/Mirwinli/coffe_plus/internal/core/transport/http/types"
	categories_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

type PatchCategoryRequest struct {
	Name core_http_types.Nullable[string] `json:"name"`
}

type PatchCategoryResponse CategoryDTOResponse

func (r *PatchCategoryRequest) Validate() error {
	if r.Name.Set && r.Name.Value == nil {
		return fmt.Errorf("name is required")
	}

	return nil
}

func (h *CategoryHTTPHandler) PatchCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request PatchCategoryRequest
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

	patch := domain.NewCategoryPatch(request.Name.ToDomain())

	params := categories_ports_in.NewPatchCategoryParams(id, patch)

	category, err := h.categoryService.PatchCategory(ctx, params)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch category",
		)
		return
	}

	response := PatchCategoryResponse{
		ID:   category.Category.ID,
		Name: category.Category.Name,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
