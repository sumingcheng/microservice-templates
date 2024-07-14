package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/config"
	"microservice/sales-system/controller"
	"microservice/sales-system/service"
)

func Category(db *gorm.DB, r *gin.RouterGroup, err *config.CustomError) {
	category := controller.Category{
		S: &service.Category{
			DB: db,
		},
		E: err,
	}

	{
		r.Group("/category").
			GET("/list/:page_size/:page_number", category.GetList).
			GET("/one/:id", category.GetOne).
			POST("/add", category.Add).
			POST("/update", category.Update).
			POST("/delete", category.Delete).
			POST("/search", category.Search)
	}
}
