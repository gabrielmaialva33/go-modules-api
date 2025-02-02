package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field         string   `json:"field"`
	Tag           string   `json:"tag"`
	Value         string   `json:"value"`
	AllowedValues []string `json:"allowed_values,omitempty"`
}

// Custom validation function for boolean pointers
func boolValidator(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() == reflect.Ptr && !field.IsNil() {
		_, ok := field.Interface().(*bool)
		return ok
	}
	return true
}

func init() {
	err := validate.RegisterValidation("is_bool", boolValidator)
	if err != nil {
		fmt.Println("Error registering custom validation:", err)
	}
}

func extractAllowedValues(param string) []string {
	if param != "" {
		return strings.Fields(param)
	}
	return nil
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

			if err.Tag() == "oneof" {
				element.AllowedValues = extractAllowedValues(err.Param())
			}

			errors = append(errors, element)
		}
	}
	return errors
}
