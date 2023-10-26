package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationsErrors := err.(validator.ValidationErrors)
	validationsError := validationsErrors[0]

	field := strings.ToLower(validationsError.StructField())
	switch validationsError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " is required with max " + validationsError.Param())
	case "min":
		return errors.New(field + " is required with min " + validationsError.Param())
	case "email":
		return errors.New(field + " is invalid")
	}

	return nil
}
