package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"users-api-server/internal/login/model"
	"users-api-server/internal/login/service"
	"users-api-server/pkg/midleware/auth"
)

type LoginHandler struct {
	loginService service.ILoginService
	jWtService   auth.IAuthService
}

func NewLoginHandler(loginService service.ILoginService, jWtService auth.IAuthService) LoginHandler {
	return LoginHandler{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller LoginHandler) Login(ctx *gin.Context) {
	var credential model.LoginCredentials
	if err := ctx.ShouldBind(&credential); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"status": "failed", "message": fmt.Sprintf("Failed to parse request body: %s", err.Error())})
		return
	}

	if isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password); isUserAuthenticated {
		ctx.JSON(http.StatusOK, gin.H{
			"token": controller.jWtService.GenerateToken(credential.Email, credential.Password),
		})
		return
	}

	ctx.AbortWithStatus(http.StatusUnauthorized)
	return
}
