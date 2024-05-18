package errs

import (
	"errors"
	"fmt"
)

func New(msg string) error {
	return errors.New(msg)
}

func WrapErrorError(errHigh error, errLow error) error {
	return fmt.Errorf("%w: %w", errHigh, errLow)
}

func WrapErrorString(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}

func WrapStringError(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
