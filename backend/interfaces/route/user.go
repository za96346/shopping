package route

import (
	"backend/interfaces/controller"
	"backend/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func Member(props *gin.RouterGroup, memberController *controller.MemberController) {
	props.Use(middleware.Permission)
	{
		props.GET("/my", memberController.GetMine)
	}
}
