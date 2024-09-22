package route

import (
	"backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)

func Member(props *gin.RouterGroup, memberController *controller.MemberController) {
	props.GET("/my", memberController.GetMine)
}
