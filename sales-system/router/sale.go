package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice/sales-system/controller"
	"microservice/sales-system/service"
	"microservice/sales-system/utils"
)

func Sale(db *gorm.DB, r *gin.RouterGroup, err *utils.CustomError) {
	sale := controller.Category{
		S: &service.Category{
			DB: db,
		},
		E: err,
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
