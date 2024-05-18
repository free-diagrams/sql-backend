package errors

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	EmptyToken   = errors.New("empty token")
	InvalidToken = errors.New("invalid token")

	Validation = errors.New("validation failed")
	Parsing    = errors.New("parsing error")

	DatabaseNotFound = errors.New("database not found")

	InternalDatabase = errors.New("internal database error")
	InternalServer   = errors.New("internal server error")
)

func ErrorToHTTPStatusAndMessageID(err error) (int, string) {
	if errors.Is(err, EmptyToken) {
		return http.StatusUnauthorized, "EmptyToken"
	} else if errors.Is(err, Parsing) {
		return http.StatusBadRequest, "UnexpectedError"
	} else if errors.Is(err, InvalidToken) {
		return http.StatusUnauthorized, "InvalidToken"
	} else if errors.Is(err, Validation) {
		return http.StatusBadRequest, "ValidationFailed"
	} else if errors.Is(err, DatabaseNotFound) {
		return http.StatusNotFound, "DatabaseNotFound"
	} else if errors.Is(err, InternalDatabase) {
		return http.StatusInternalServerError, "InternalDatabaseError"
	} else if errors.Is(err, InternalServer) {
		return http.StatusInternalServerError, "UnexpectedError"
	} else {
		return http.StatusInternalServerError, "UnexpectedError"
	}
}
