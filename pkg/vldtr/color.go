package vldtr

import (
	"github.com/free-diagrams/sql-backend/pkg/errs"
	"github.com/pkg/errors"
	"regexp"
)

func RGB(value string) error {
	match, err := regexp.MatchString(`^#[0-9a-fA-F]{6}$`, value)
	if err != nil {
		return errors.Wrap(err, "failed to check rgb code")
	}
	if !match {
		return errors.Wrap(errs.Validation, "rgb color code contains invalid characters")
	}

	return nil
}
