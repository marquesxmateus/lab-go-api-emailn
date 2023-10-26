package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New("Internal Server Error")
var ErrNotFound error = errors.New("Not Found")

func ProcessErrorToReturn(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternal
	}
	return err
}
