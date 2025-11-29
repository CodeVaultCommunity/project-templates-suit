// Package errorsapp provides utilities for handling and representing
// application-level errors in a structured and consistent way.
package errorsapp

// AppError represents a structured application error.
//
// It provides fields for an HTTP status code, a machine-readable error code,
// a human-readable message, and an optional underlying error.
// Fields tagged with `json:"-"` are not serialized when returning the error as JSON.
type AppError struct {
	Status  int    `json:"-"`              // HTTP status code (not serialized)
	Code    string `json:"code,omitempty"` // Optional machine-readable error code
	Message string `json:"message"`        // Human-readable error message
	Err     error  `json:"-"`              // Underlying error (not serialized)
}

// Error implements the built-in error interface.
//
// It returns the underlying error message if present,
// otherwise it returns the AppError's message.
func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

// New creates and returns a new AppError instance.
//
// Parameters:
//   - status: HTTP status code associated with the error.
//   - code: machine-readable code to identify the error type.
//   - message: human-readable description of the error.
//   - err: optional underlying error for debugging/logging.
//
// Example:
//
//	err := errorsapp.New(http.StatusBadRequest, "INVALID_INPUT", "Invalid user data", nil)
func New(status int, code string, message string, err error) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
		Err:     err,
	}
}
