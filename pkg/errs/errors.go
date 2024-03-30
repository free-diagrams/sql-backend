package errs

import "github.com/pkg/errors"

var (
	InternalDatabase = errors.New("internal database error")
)
