package validate

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

func ValidPhone(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(fl.Field().String())
}

func RegisterValidations(v *validator.Validate) error {
	validations := map[string]validator.Func{
		"is-phone": ValidPhone,
	}

	for tag, fn := range validations {
		if err := v.RegisterValidation(tag, fn); err != nil {
			zap.S().Errorf("Failed to register validation for '%s': %v", tag, err)
			return err
		}
	}
	return nil
}

func RegisterValTrans(v *validator.Validate, trans ut.Translator) error {
	if err := RegisterValidations(v); err != nil {
		return err
	}
	if err := v.RegisterTranslation("is-phone", trans, regFunc, transFunc); err != nil {
		return err
	}
	return nil
}

func regFunc(ut ut.Translator) error {
	return ut.Add("is-phone", "手机号格式不正确", true)
}

func transFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("is-phone", fe.Field())
	return t
}
