package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct checks for errors and returns a simple string error
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errorMsgs []string
		for _, e := range validationErrors {
			msg := fmt.Sprintf("field '%s' failed validation: %s", e.Field(), e.Tag())
			errorMsgs = append(errorMsgs, msg)
		}
		return fmt.Errorf("validation error: %s", strings.Join(errorMsgs, ", "))
	}

	return err
}
