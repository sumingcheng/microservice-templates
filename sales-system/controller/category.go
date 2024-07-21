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

func (ce *Category) GetList(c *gin.Context) {
	pageSize, pnErr := strconv.Atoi(c.Param("page_size"))
	if pnErr != nil {
		c.JSON(http.StatusBadRequest, ce.E.BadParameter(pnErr))
		return
	}
	pageNumber, psErr := strconv.Atoi(c.Param("page_number"))
	if psErr != nil {
		c.JSON(http.StatusBadRequest, ce.E.BadParameter(psErr))
		return
	}

	data, err := ce.S.GetList(pageSize, pageNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ce.E.QueryDataFailed(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func (ce *Category) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ce.E.BadParameter(err))
		return
	}

	data, err := ce.S.GetOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ce.E.QueryDataFailed(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}

func (ce *Category) Add(c *gin.Context) {
	var body CateAddBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ce.E.BadParameter(err))
		return
	}

	// 检查分类名是否已存在
	exists, err := ce.S.Exists(body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ce.E.BadParameter(err))
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "分类名已存在",
		})
		return
	}

	id, err := ce.S.Add(body.Name)
	if err != nil || id == 0 {
		c.JSON(http.StatusInternalServerError, ce.E.CreateDataFailed(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": id,
	})
}

type CateUpdateBody struct {
	Id       int    `json:"id" binding:"required"`
	CateName string `json:"cate_name" binding:"required"`
}

func (ce *Category) Update(c *gin.Context) {
	var body CateUpdateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, ce.E.BadParameter(err))
		return
	}

	id, err := ce.S.Update(int32(body.Id), body.CateName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ce.E.UpdateDataFailed(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": id,
	})
}

func (ce *Category) Delete(c *gin.Context) {

}

func (ce *Category) Search(c *gin.Context) {

}

func (ce *Category) SearchWitchKeyWord(c *gin.Context) {

}

func (ce *Category) SearchWitchDate(c *gin.Context) {

}
