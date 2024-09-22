package controller

import (
	"backend/application/service"
	"backend/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
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
			"data":    *memberEntity,
		},
	)
}

func (m *MemberController) Login(Request *gin.Context) {
	session := sessions.Default(Request)

	// 請求處理
	reqBody := new(struct {
		Password string
		Account  string
	})

	if Request.ShouldBindJSON(&reqBody) != nil {
		Request.JSON(
			http.StatusOK,
			gin.H{
				"message": "Request Data 格式不正確",
			},
		)
		return
	}

	memberInfo, err := m.memberApp.Login(&entities.Member{
		Account: reqBody.Account,
		Password: reqBody.Password,
	})

	if err != nil || memberInfo.ID == 0 {
		Request.JSON(
			http.StatusOK,
			gin.H{
				"message": "登入失敗",
			},
		)
		return
	}

	session.Set("isLogin", "Y")
	session.Set("memberId", strconv.Itoa(memberInfo.ID))

	session.Save()

	Request.JSON(
		http.StatusOK,
		gin.H{
			"message": "登入成功",
			"data": *memberInfo,
		},
	)
}