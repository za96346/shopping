package controller

import (
	"backend/application/service"
	"backend/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderApp service.OrderAppInterface
}

func NewOrder(app service.OrderAppInterface) *OrderController {
	return &OrderController{
		orderApp: app,
	}
}

func (o *OrderController) GetOrders(Request *gin.Context) {
	orderEntity, err := o.orderApp.GetOrders(&entities.Order{})

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
			"data":    *orderEntity,
		},
	)
}

func (o *OrderController) AddOrder(Request *gin.Context) {
	o.orderApp.AddOrder(&entities.Order{})
}