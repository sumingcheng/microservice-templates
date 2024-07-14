package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/config"
	"microservice/sales-system/router"
)

func Router(db *gorm.DB, r *gin.Engine, err *config.CustomError) gin.HandlerFunc {
	apiRouter := r.Group("/v1")

	return func(context *gin.Context) {
		router.Category(db, apiRouter, err)
		router.Product(db, apiRouter, err)
		router.Sale(db, apiRouter, err)
	}
}
