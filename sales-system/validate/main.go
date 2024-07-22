package validate

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func RegisterCustomValidations(v *validator.Validate) {
	validations := map[string]validator.Func{
		"is-phone":    IsValidPhone,
		"is-password": IsPasswordValid,
	}

	for tag, fn := range validations {
		if err := RegisterValidation(v, tag, fn); err != nil {
			zap.S().Errorf("Error registering validation for %s: %v", tag, err)
		}
	}
}
