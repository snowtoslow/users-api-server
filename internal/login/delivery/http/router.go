package http

import (
	"github.com/gin-gonic/gin"
	"users-api-server/internal/login/service"
	"users-api-server/pkg/midleware/auth"
)

func RegisterHttpEndPoints(
	router *gin.RouterGroup,
	loginService service.ILoginService,
	jWtService auth.IAuthService,
) {
	h := NewLoginHandler(loginService, jWtService)

	userRouter := router.Group("/auth")
	{
		userRouter.POST("", h.Login)
	}
}
