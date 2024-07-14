package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

var Translate ut.Translator

func CreateTranslator(locale string) error {
	validate, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		fmt.Println("类型断言失败") // 添加了打印语句
		return fmt.Errorf("类型断言失败")
	}

	// 注册结构体的json标签名作为字段名
	validate.RegisterTagNameFunc(func(sf reflect.StructField) string {
		name := sf.Tag.Get("json")
		if name == "_" {
			return ""
		}
		return name
	})

	zhTranslation := zh.New()
	enTranslation := en.New()
	uniTranslations := ut.New(enTranslation, zhTranslation, enTranslation)

	translator, ok := uniTranslations.GetTranslator(locale)
	if !ok {
		fmt.Printf("找不到翻译器(%s)\n", locale)
		return fmt.Errorf("找不到翻译器(%s)", locale)
	}

	switch locale {
	case "zh":
		if err := zhtranslations.RegisterDefaultTranslations(validate, translator); err != nil {
			zap.S().Errorf("注册中文翻译器失败: %v", err)
			return fmt.Errorf("注册中文翻译器失败: %v", err)
		}
	case "en":
		if err := entranslations.RegisterDefaultTranslations(validate, translator); err != nil {
			zap.S().Errorf("注册英文翻译器失败: %v", err)
			return fmt.Errorf("注册英文翻译器失败: %v", err)
		}
	default:
		zap.S().Errorf("不支持的语言类型(%s)", locale)
		return fmt.Errorf("不支持的语言类型(%s)", locale)
	}

	Translate = translator // 设置全局翻译器
	return nil
}

func RemoveStructName(errs validator.ValidationErrors) map[string]string {
	errors := map[string]string{}
	for _, e := range errs {
		field := strings.SplitN(e.Namespace(), ".", -1)
		fieldName := field[len(field)-1]
		errors[fieldName] = e.Translate(Translate)
	}
	return errors
}
