package vldtr

import (
	"errors"
)

func NotEmptyString(value string) error {
	if len([]rune(value)) == 0 {
		return errors.New("string is empty")
	}
	return nil
}

func NotEmptyStringPtr(value *string) error {
	if value == nil {
		return errors.New("pointer does not point to string")
	}

	return NotEmptyString(*value)
}
