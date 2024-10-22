package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Invalid Field: '%s', Condition: '%s'", e.Field(), e.Value())
			errorMessages = append(errorMessages, errorMessage)
		}
		return errors.New(strings.Join(errorMessages, " | "))
	}
	return nil
}
