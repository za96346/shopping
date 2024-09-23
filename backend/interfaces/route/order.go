package route

import (
	"backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)

func Order(props *gin.RouterGroup, orderController *controller.OrderController) {
	props.GET("/", orderController.GetOrders)
	props.PUT("/", orderController.AddOrder)
}
