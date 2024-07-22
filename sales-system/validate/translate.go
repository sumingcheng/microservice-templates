package validate

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

var trans ut.Translator

func TransInit(locale string) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)

		if !ok {
			zap.S().Fatalf("uni.GetTranslator(%s) failed", locale)
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册自定义验证规则及其翻译
		if err := RegisterValTrans(v, trans); err != nil {
			zap.S().Errorf("Failed to register validation translations: %v", err)
			return err
		}

		// 使用 JSON 标签作为字段名称
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0] // 获取 json 标签的第一部分
			if name == "-" {
				return "" // 如果标签是 "-"，则不使用任何名称
			}
			return name
		})

		switch locale {
		case "zh":
			return chTranslations.RegisterDefaultTranslations(v, trans)
		case "en":
			return enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			return enTranslations.RegisterDefaultTranslations(v, trans)
		}
	}

	zap.S().Error("Failed to assert Validator")
	return fmt.Errorf("failed to assert Validator")
}

func TranslateErrors(err error) string {
	if err == nil {
		return ""
	}

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var errMessages []string
		for _, e := range errs {
			translatedMsg := e.Translate(trans)
			errMessages = append(errMessages, translatedMsg)
		}

		return strings.Join(errMessages, "; ")
	}
	return err.Error()
}
