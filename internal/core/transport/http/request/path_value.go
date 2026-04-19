package core_http_request

import (
	"fmt"
	"net/http"

	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	"github.com/google/uuid"
)

func GetUUIDPathValue(r *http.Request, key string) (uuid.UUID, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return uuid.Nil, fmt.Errorf(
			"no key=`%s` in path values: %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	ID, err := uuid.Parse(pathValue)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf(
			"path value=`%s` by key=`%s` not a valid uuif: %v: %w ",
			pathValue,
			key,

			core_errors.ErrInvalidArgument,
		)
	}

	return ID, nil
}
