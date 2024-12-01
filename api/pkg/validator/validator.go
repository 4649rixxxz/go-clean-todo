package validator

import (
	"go-clean-todo/presentation/settings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}

	return validate
}

func MakeValidationErrMessages(err error) []settings.ErrorResponse {
	var errResponse []settings.ErrorResponse
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errResponse = append(errResponse, settings.ErrorResponse{
				Field:   fieldError.Field(),
				Message: fieldError.Error(),
			})
		}
	}

	return errResponse
}
