package controller

import (
	"github.com/gin-gonic/gin"
	"microservice/sales-system/service"
	"microservice/sales-system/utils"
	"net/http"
)

type Category struct {
	S *service.Category
	E *utils.CustomError
}

type CateAddBody struct {
	Name string `json:"name" binding:"required"`
}

func (ca *Category) GetList(c *gin.Context) {

}

func (ca *Category) GetOne(c *gin.Context) {

}

func (ca *Category) Add(c *gin.Context) {
	var body CateAddBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ca.E.BadParameter(err))
		return
	}

	id, err := ca.S.Add(body.Name)
	if err != nil || id == 0 {
		c.JSON(http.StatusInternalServerError, ca.E.CreateDataFailed(err))
		return
	}

	c.JSON(http.StatusOK, ca.E.Success())
}

func (ca *Category) Update(c *gin.Context) {

}

func (ca *Category) Delete(c *gin.Context) {

}

func (ca *Category) Search(c *gin.Context) {

}

func (ca *Category) SearchWitchKeyWord(c *gin.Context) {

}

func (ca *Category) SearchWitchDate(c *gin.Context) {

}
