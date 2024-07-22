package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SetValidate() {
	// 初始化翻译器
	if err := TransInit("zh"); err != nil {
		zap.S().Fatalf("Failed to initialize translator: %v", err)
	}

	// 注册自定义验证规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := RegisterValidations(v)
		if err != nil {
			zap.S().Fatalf("Failed to register custom validations: %v", err.Error())
			return
		}
	}
}
