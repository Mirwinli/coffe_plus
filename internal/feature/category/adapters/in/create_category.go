package category_adapters_in

import (
	"net/http"

	domain2 "github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_request "github.com/Mirwinli/coffe_plus/internal/core/transport/http/request"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	category_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/category/ports/in"
)

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100" example:"Drinks"`
}

type CreateCategoryResponse CategoryDTOResponse

// CreateCategory godoc
// @Summary Створення категорії
// @Description Створення категорії продуктів
// @Description Only for admins
// @Tags category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateCategoryRequest true "CreateCategory тіло запиту"
// @Success 201 {object} CreateCategoryRequest "Категорія"
// @Failure 400 {object} core_http_response.ErrorResponse.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse.ErrorResponse "Unauthorized"
// @Failure 500 {object} core_http_response.ErrorResponse.ErrorResponse "Internal server error"
// @Router /category [post]
func (h *CategoryHTTPHandler) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateCategoryRequest
	if err := core_http_request.DecodeAndValidate(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	domain := domain2.NewCategoryUninitialized(request.Name)

	in := category_ports_in.NewCreateCategoryParams(domain)
	result, err := h.categoryService.CreateCategory(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create category",
		)
		return
	}

	response := CreateCategoryResponse(categoryDTOFromDomain(result.Category))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
