package vldtr

import (
	"github.com/free-diagrams/sql-backend/pkg/errs"
	"github.com/pkg/errors"
)

func NotEmptyString(value string) error {
	if len([]rune(value)) == 0 {
		return errors.Wrap(errs.Validation, "empty string")
	}
	return nil
}

func NotEmptyStringPtr(value *string) error {
	if value == nil {
		return errors.Wrap(errs.Validation, "pointer does not point to string")
	}

	return NotEmptyString(*value)
}
