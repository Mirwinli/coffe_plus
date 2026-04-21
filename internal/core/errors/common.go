package core_errors

import "errors"

var (
	ErrConflict             = errors.New("conflict")
	ErrForbidden            = errors.New("forbidden")
	ErrNotFound             = errors.New("not found")
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrInternalServerError  = errors.New("internal server error")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrForeignKeyViolation  = errors.New("foreign key violation")
	ErrUniqueViolation      = errors.New("unique violation")
	ErrProductIsntAvailable = errors.New("product is not available")
)
