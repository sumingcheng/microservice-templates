package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/controller"
	"microservice/sales-system/service"
)

func Sale(db *gorm.DB, r *gin.RouterGroup) {
	sale := controller.Category{
		S: &service.Category{
			DB: db,
		},
	}

	{
		r.Group("/sale").
			GET("/list/:page_size/:page_number", sale.GetList).
			GET("/one/:id", sale.GetOne).
			POST("/add", sale.Add).
			POST("/update", sale.Update).
			POST("/delete", sale.Delete).
			POST("/search", sale.SearchWitchKeyWord).
			POST("/search/date", sale.SearchWitchDate)
	}
}
