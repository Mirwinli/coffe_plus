package products_adapters_in_products_transport_http

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_response "github.com/Mirwinli/coffe_plus/internal/core/transport/http/response"
	products_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/products/ports/in"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateProductRequest struct {
	Name        string
	Description *string
	Price       domain.Money
	IsAvailable *bool
	CategoryID  uuid.UUID
	ImageFile   multipart.File
	ImageHeader *multipart.FileHeader
}

type CreateProductResponse ProductDTOResponse

func (r *CreateProductRequest) Validate() error {
	if r.Price.IsZero() {
		return fmt.Errorf(
			"price must be not zero: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if r.CategoryID == uuid.Nil {
		return fmt.Errorf(
			"Category is required: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	nameLen := len([]rune(r.Name))
	if nameLen < 1 || nameLen > 100 {
		return fmt.Errorf(
			"name must be between 1 and 100 characters: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if r.Description != nil {
		descriptionLen := len([]rune(*r.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"description must be between 1 and 1000 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if r.IsAvailable == nil {
		return fmt.Errorf(
			"isAvailable must be set: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	ext := filepath.Ext(r.ImageHeader.Filename)
	if !allowedFormatImage[strings.ToLower(ext)] {
		return fmt.Errorf(
			"image must be jpg, jpeg,webp or png: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if r.ImageHeader.Size > 5<<20 {
		return fmt.Errorf(
			"image size must be less than 5MB: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (h *ProductsHTTPHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	request, err := Parse(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to parse request",
		)
		return
	}
	defer request.ImageFile.Close()

	if err = request.Validate(); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to validate request",
		)
		return
	}

	in := products_ports_in.NewCreateProductParams(
		request.Name,
		request.Description,
		request.Price,
		*request.IsAvailable,
		request.CategoryID,
		request.ImageFile,
		request.ImageHeader.Filename,
	)

	result, err := h.productsService.CreateProduct(ctx, in)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create product",
		)
		return
	}

	response := GetProductResponse(productDTOFromDomain(result.Product))

	responseHandler.JSONResponse(
		response,
		http.StatusCreated,
	)
}

func Parse(r *http.Request) (CreateProductRequest, error) {
	name := r.FormValue("name")

	var descriptionPtr *string
	if description := r.FormValue("description"); description != "" {
		descriptionPtr = &description
	}

	categoryID, err := uuid.Parse(r.FormValue("category_id"))
	if err != nil {
		return CreateProductRequest{}, fmt.Errorf(
			"invalid category_id: %w",
		)
	}

	isAvailable, err := strconv.ParseBool(r.FormValue("is_available"))
	if err != nil {
		return CreateProductRequest{}, fmt.Errorf(
			"convert string to bool: %w", err,
		)
	}

	priceStr := r.FormValue("price")
	price, err := decimal.NewFromString(priceStr)
	if err != nil {
		return CreateProductRequest{}, fmt.Errorf(
			"invalid price: %w",
		)
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		return CreateProductRequest{}, fmt.Errorf(
			"get image: %w", err,
		)
	}

	request := CreateProductRequest{
		Name:        name,
		Description: descriptionPtr,
		Price:       price,
		IsAvailable: &isAvailable,
		CategoryID:  categoryID,
		ImageFile:   file,
		ImageHeader: header,
	}

	return request, nil
}
