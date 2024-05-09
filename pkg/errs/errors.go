package errs

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	EmptyTokenText   = "empty token"
	EmptyToken       = errors.New(EmptyTokenText)
	InvalidTokenText = "invalid token"
	InvalidToken     = errors.New(InvalidTokenText)

	ValidationText = "validation failed"
	Validation     = errors.New(ValidationText)

	DatabaseNotFoundText = "database not found"
	DatabaseNotFound     = errors.New(DatabaseNotFoundText)

	InternalDatabaseText = "internal database error"
	InternalDatabase     = errors.New(InternalDatabaseText)
)

func ErrorToHTTPStatusAndMessageID(err error) (int, string) {
	if errors.Is(err, EmptyToken) {
		return http.StatusUnauthorized, "EmptyToken"
	} else if errors.Is(err, InvalidToken) {
		return http.StatusUnauthorized, "InvalidToken"
	} else if errors.Is(err, Validation) {
		return http.StatusBadRequest, "ValidationFailed"
	} else if errors.Is(err, DatabaseNotFound) {
		return http.StatusNotFound, "DatabaseNotFound"
	} else if errors.Is(err, InternalDatabase) {
		return http.StatusInternalServerError, "InternalDatabaseError"
	} else {
		return http.StatusInternalServerError, "UnspecifiedError"
	}
}
