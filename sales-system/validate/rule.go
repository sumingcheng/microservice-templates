package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	PhoneRegex  = regexp.MustCompile(`^1[3-9]\d{9}$`)
	NumberRegex = regexp.MustCompile(`[0-9]`)
	LetterRegex = regexp.MustCompile(`[a-zA-Z]`)
)

func IsValidPhone(fl validator.FieldLevel) bool {
	return PhoneRegex.MatchString(fl.Field().String())
}

func IsPasswordValid(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return NumberRegex.MatchString(password) && LetterRegex.MatchString(password)
}

func RegisterValidation(v *validator.Validate, tag string, fn validator.Func) error {
	return v.RegisterValidation(tag, fn)
}
