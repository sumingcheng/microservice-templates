package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/controller"
	"microservice/sales-system/service"
)

func Product(db *gorm.DB, r *gin.RouterGroup) {
	product := controller.Category{
		S: &service.Category{
			DB: db,
		},
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
