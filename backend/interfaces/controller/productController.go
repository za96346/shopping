package controller

import (
	"backend/application/service"
	"backend/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productApp service.ProductAppInterface
}

func NewProduct(app service.ProductAppInterface) *ProductController {
	return &ProductController{
		productApp: app,
	}
}

func (p *ProductController) GetProducts(Request *gin.Context) {
	productEntity, err := p.productApp.GetProducts(&entities.Product{

	})

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
			"data":    *productEntity,
		},
	)
}
