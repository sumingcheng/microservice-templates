package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

func Paginate(page, pageSize int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

func HandleBinError(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  RemoveStructName(errs),
		})

		return
	}

	// 如果错误不是验证错误，返回通用错误消息
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  err.Error(),
	})
}
