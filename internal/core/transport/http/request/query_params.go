package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
)

func GetOffsetLimitQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParam  = "limit"
		offsetQueryParam = "offset"
	)

	limit, err := GetIntQueryParams(r, limitQueryParam)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get `offset` query param: %w",
			err,
		)
	}

	offset, err := GetIntQueryParams(r, offsetQueryParam)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"get `offset` query param: %w",
			err,
		)
	}

	return offset, limit, nil
}

func GetIntQueryParams(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param=%s bt key=%s not valid for integer: %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &val, nil
}

func GetUUIDQueryParams(r *http.Request, key string) (*uuid.UUID, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := uuid.Parse(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param=%s bY key=%s not valid for uuid: %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &val, nil
}

func GetStringQueryParams(r *http.Request, key string) *string {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil
	}

	return &param
}
