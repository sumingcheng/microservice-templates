package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/utils"
)

func SetRouter(db *gorm.DB, r *gin.Engine) {
	apiRouter := r.Group("/v1")
	Category(db, apiRouter, &utils.CustomError{})
	Product(db, apiRouter, &utils.CustomError{})
	Sale(db, apiRouter, &utils.CustomError{})
}
