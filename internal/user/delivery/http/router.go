package http

import (
	"github.com/gin-gonic/gin"
	"users-api-server/internal/user"
)

func RegisterHttpEndPoints(router *gin.RouterGroup, userService user.IUserService) {
	h := NewHandler(userService)

	userRouter := router.Group("/users")
	{
		userRouter.GET("", h.GetAll)
		userRouter.GET("/:id", h.GetUserByID)
		userRouter.POST("", h.CreateUser)
		userRouter.PUT("/:id", h.UpdateUser)
		userRouter.PATCH("/:id", h.SetPartially)
	}
}
