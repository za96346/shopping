package controller

import (
	"backend/application/service"
	"backend/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryApp service.CategoryAppInterface
}

func NewCategory(app service.CategoryAppInterface) *CategoryController {
	return &CategoryController{
		categoryApp: app,
	}
}

func (c *CategoryController) GetCategories(Request *gin.Context) {
	categoryEntity, err := c.categoryApp.GetCategories(&entities.Category{})

	if err != nil {
		Request.JSON(
			http.StatusForbidden,
			gin.H{
				"message": "失敗",
				"data":    nil,
			},
		)
	}

	Request.JSON(
		http.StatusOK,
		gin.H{
			"message": "成功",
			"data":    *categoryEntity,
		},
	)
}
