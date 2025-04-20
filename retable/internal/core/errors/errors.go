
package errors

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
