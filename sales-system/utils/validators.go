package utils

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"regexp"
)

var (
	phoneRegex  = regexp.MustCompile(`^1[3-9]\d{9}$`)
	numberRegex = regexp.MustCompile(`[0-9]`)
	letterRegex = regexp.MustCompile(`[a-zA-Z]`)
)

func IsValidPhone(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(fl.Field().String())
}

func IsPasswordValid(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return numberRegex.MatchString(password) && letterRegex.MatchString(password)
}

func registerValidation(v *validator.Validate, tag string, fn validator.Func) {
	if err := v.RegisterValidation(tag, fn); err != nil {
		zap.S().Fatalf("Error registering validation for %s: %v", tag, err)
	}
}

func RegisterCustomValidations(v *validator.Validate) {
	validations := map[string]validator.Func{
		"is-valid-phone":    IsValidPhone,
		"is-password-valid": IsPasswordValid,
	}

	for tag, fn := range validations {
		registerValidation(v, tag, fn)
	}
}
