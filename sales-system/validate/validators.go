package validate

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

func registerValidation(v *validator.Validate, tag string, fn validator.Func) error {
	return v.RegisterValidation(tag, fn)
}

func RegisterCustomValidations(v *validator.Validate) {
	validations := map[string]validator.Func{
		"is-phone":    IsValidPhone,
		"is-password": IsPasswordValid,
	}

	for tag, fn := range validations {
		if err := registerValidation(v, tag, fn); err != nil {
			zap.S().Errorf("Error registering validation for %s: %v", tag, err)
		}
	}
}
