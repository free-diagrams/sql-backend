package vldtr

import (
	"errors"
	"regexp"
)

func RGB(value string) error {
	match, err := regexp.MatchString(`^#[0-9a-fA-F]{6}$`, value)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("rgb color code contains invalid characters")
	}

	return nil
}
