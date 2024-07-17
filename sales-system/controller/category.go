package controller

import (
	"github.com/gin-gonic/gin"
	"microservice/sales-system/service"
	"microservice/sales-system/utils"
	"net/http"
	"strconv"
)

type Category struct {
	S *service.Category
	E *utils.CustomError
}

type CateAddBody struct {
	Name string `json:"name" binding:"required"`
}

func (ca *Category) GetList(c *gin.Context) {
	pageSize, pnErr := strconv.Atoi(c.Param("page_size"))
	if pnErr != nil {
		c.JSON(http.StatusBadRequest, ca.E.BadParameter(pnErr))
		return
	}
	pageNumber, psErr := strconv.Atoi(c.Param("page_number"))
	if psErr != nil {
		c.JSON(http.StatusBadRequest, ca.E.BadParameter(psErr))
		return
	}

	data, err := ca.S.GetList(pageSize, pageNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ca.E.QueryDataFailed(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"name": "success",
		"code": 0,
		"msg":  "ok",
		"err":  utils.TranslateErrors(err),
	})
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
