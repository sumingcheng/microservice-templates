package controller

import (
	"github.com/gin-gonic/gin"
	"microservice/sales-system/config"
	"microservice/sales-system/service"
)

type Product struct {
	S *service.Product
	E *config.CustomError
}

func (pr *Product) GetList(c *gin.Context) {

}

func (pr *Product) GetOne(c *gin.Context) {

}

func (pr *Product) Add(c *gin.Context) {

}

func (pr *Product) Update(c *gin.Context) {

}

func (pr *Product) Delete(c *gin.Context) {

}

func (pr *Product) Search(c *gin.Context) {

}
