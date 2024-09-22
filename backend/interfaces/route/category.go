package route

import (
	"backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)

func Category(props *gin.RouterGroup, categoryController *controller.CategoryController) {
	props.GET("/", categoryController.GetCategories)
}
