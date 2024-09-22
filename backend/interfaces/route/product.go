package route

import (
	"backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)

func Product(props *gin.RouterGroup, productController *controller.ProductController) {
	props.GET("/", productController.GetProducts)
}
