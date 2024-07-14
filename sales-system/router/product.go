package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/config"
	"microservice/sales-system/controller"
	"microservice/sales-system/service"
)

func Product(db *gorm.DB, r *gin.RouterGroup, err *config.CustomError) {
	product := controller.Category{
		S: &service.Category{
			DB: db,
		},
		E: err,
	}

	{
		r.Group("/product").
			GET("/list/:page_size/:page_number", product.GetList).
			GET("/one/:id", product.GetOne).
			POST("/add", product.Add).
			POST("/update", product.Update).
			POST("/delete", product.Delete).
			POST("/search", product.Search)
	}
}
