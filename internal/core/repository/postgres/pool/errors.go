package core_postgres_pool

import "errors"

var (
	ErrNoRows         = errors.New("no rows")
	ErrViolatesUnique = errors.New("unique_violation")
	ErrUnknown        = errors.New("unknown error")
)
