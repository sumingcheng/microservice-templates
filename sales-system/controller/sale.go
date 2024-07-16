package controller

import (
	"github.com/gin-gonic/gin"
	"microservice/sales-system/service"
	"microservice/sales-system/utils"
)

type Sale struct {
	S *service.Sale
	E *utils.CustomError
}

func (sa *Sale) GetList(c *gin.Context) {
}

func (sa *Sale) GetOne(c *gin.Context) {

}

func (sa *Sale) Add(c *gin.Context) {

}

func (sa *Sale) Update(c *gin.Context) {

}

func (sa *Sale) Delete(c *gin.Context) {

}

func (sa *Sale) SearchWitchKeyWord(c *gin.Context) {

}

func (sa *Sale) SearchWitchDate(c *gin.Context) {

}
