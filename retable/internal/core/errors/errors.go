
package errors

import "fmt"

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewAppError(code string, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	ErrNotFound        = NewAppError("NOT_FOUND", "Resource not found", nil)
	ErrUnauthorized    = NewAppError("UNAUTHORIZED", "Unauthorized access", nil)
	ErrForbidden       = NewAppError("FORBIDDEN", "Access forbidden", nil)
	ErrInvalidInput    = NewAppError("INVALID_INPUT", "Invalid input data", nil)
	ErrDuplicateEntry  = NewAppError("DUPLICATE_ENTRY", "Resource already exists", nil)
	ErrInternalServer  = NewAppError("INTERNAL_SERVER", "Internal server error", nil)
)

type AppError struct {
	Code    string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewAppError(code, message string) AppError {
	return AppError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrNotFound = NewAppError("NOT_FOUND", "Resource not found")
	ErrInvalidInput = NewAppError("INVALID_INPUT", "Invalid input provided")
	ErrUnauthorized = NewAppError("UNAUTHORIZED", "Unauthorized access")
	ErrForbidden = NewAppError("FORBIDDEN", "Forbidden access")
)
