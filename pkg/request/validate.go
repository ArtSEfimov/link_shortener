package request

import "github.com/go-playground/validator/v10"

func IsValid[T any](payload T) error {
	loginValidator := validator.New()
	validateErr := loginValidator.Struct(&payload)
	return validateErr
}
