package cart_adapters_out_redis

import (
	"fmt"
	"strconv"
	"time"
)

const (
	fieldCreatedAt = "created_at"
	fieldUpdatedAt = "updated_at"
	dayExpire      = 24 * time.Hour
)

func getCreatedUpdatedAt(values map[string]string) (time.Time, time.Time, error) {
	created, err := strconv.ParseInt(values[fieldCreatedAt], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf(
			"convert createdAt to integer: %w",
			err,
		)
	}

	updated, err := strconv.ParseInt(values[fieldUpdatedAt], 10, 64)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf(
			"convert updatedAt to integer: %w",
			err,
		)
	}

	return time.Unix(created, 0), time.Unix(updated, 0), nil
}
