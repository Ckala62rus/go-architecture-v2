package mocks

import "errors"

// NotFoundError represents a not found error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// NewNotFoundError creates a new NotFoundError
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// NewValidationError creates a new ValidationError
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{Field: field, Message: message}
}

// DatabaseError represents a database error
type DatabaseError struct {
	Message string
}

func (e *DatabaseError) Error() string {
	return "Database error: " + e.Message
}

// NewDatabaseError creates a new DatabaseError
func NewDatabaseError(message string) *DatabaseError {
	return &DatabaseError{Message: message}
}

// Common errors
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
)
