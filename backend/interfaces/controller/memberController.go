package controller

import (
	"backend/application/service"
	"backend/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
	memberApp service.MemberAppInterface
}

func NewMember(app service.MemberAppInterface) *MemberController {
	return &MemberController{
		memberApp: app,
	}
}

func (m *MemberController) GetMine(Request *gin.Context) {
	memberEntity, err := m.memberApp.GetMember(&entities.Member{})

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
			"data":    *&memberEntity,
		},
	)
}
