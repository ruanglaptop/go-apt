package httpError

import "net/http"

// ForbiddenError is a struct to contain forbidden http error
type ForbiddenError struct {
	BaseError
}

// Forbidden is a constructor to create ForbiddenError instance
func Forbidden(err error) *ForbiddenError {
	return &ForbiddenError{
		BaseError: code(http.StatusForbidden, err),
	}
}