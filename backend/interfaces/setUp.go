package interfaces

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"

	application "backend/application/service"
	"backend/infrastructure/persistence"
	"backend/interfaces/controller"
	"backend/interfaces/middleware"
	"backend/interfaces/route"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
)

// 設定 http server
func SetUp(repo *persistence.Repositories) *gin.Engine {
	port := os.Getenv("PORT")
	apiServer := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := memstore.NewStore(
		[]byte("secret11111"),
	)

	// 添加全域middleware
	apiServer.Use(
		middleware.CORS(),
		sessions.Sessions("workapp_session", store),
		middleware.RateLimit(time.Second, 100, 100),
	)

	// 新增 route group
	memberApi := apiServer.Group("/workApp/member")

	// 實例 app
	memberController := controller.NewMember(
		&application.MemberApp{
			MemberRepository: repo.Member,
		},
	)

	// 嵌入 route group
	route.Member(memberApi, memberController)
	// start
	apiServer.Run(":" + port)

	fmt.Print("api server started successfully.")

	return apiServer
}