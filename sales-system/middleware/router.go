package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/router"
)

func Router(db *gorm.DB, r *gin.Engine) gin.HandlerFunc {
	apiRouter := r.Group("/v1")

	return func(context *gin.Context) {
		router.Category(db, apiRouter)
		router.Product(db, apiRouter)
		router.Sale(db, apiRouter)
	}
}
