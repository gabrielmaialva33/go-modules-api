package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateStruct(data interface{}) []*ValidationError {
	var errors []*ValidationError
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := &ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: fmt.Sprintf("%v", err.Value()),
			}
			errors = append(errors, element)
		}
	}
	return errors
}
