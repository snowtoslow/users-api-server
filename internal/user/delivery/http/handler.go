package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"users-api-server/internal/user"
	"users-api-server/internal/user/model"
	"users-api-server/pkg/pagination"
)

type Handler struct {
	service user.IUserService
}

func NewHandler(service user.IUserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler Handler) CreateUser(ctx *gin.Context) {
	var userReqModel model.UserReq
	if err := ctx.Bind(&userReqModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"status": "failed", "message": fmt.Sprintf("Failed to parse request body: %s", err.Error())})
	}

	createdUserID, err := handler.service.CreateUser(ctx.Request.Context(), userReqModel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUserID)
}

func (handler Handler) UpdateUser(ctx *gin.Context) {
	var userReqModel model.UserReq
	if err := ctx.Bind(&userReqModel); err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"status": "failed", "message": fmt.Sprintf("Failed to parse request body: %s", err.Error())})
		return
	}

	if id := ctx.Param("id"); id != "" {
		updatedUser, err := handler.service.UpdateUserByID(ctx.Request.Context(), id, userReqModel)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, updatedUser)
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed",
		"message": fmt.Sprintf("And invalid or empty id was specified")})

}

func (handler Handler) GetUserByID(ctx *gin.Context) {
	if id := ctx.Param("id"); id != "" {
		foundUser, err := handler.service.GetUserByID(ctx.Request.Context(), id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"status": "failed",
					"message": fmt.Sprintf("there is an error which occured trying to get user: %s", err.Error())})
			return
		}
		ctx.JSON(http.StatusOK, foundUser)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed",
		"message": fmt.Sprintf("And invalid or empty id was specified")})
}

func (handler Handler) SetPartially(ctx *gin.Context) {
	partiallyFieldsToBeSet := make(map[string]interface{})
	if err := ctx.Bind(&partiallyFieldsToBeSet); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"status": "failed", "message": fmt.Sprintf("Failed to parse request body: %s", err.Error())})
		return
	}

	if id := ctx.Param("id"); id != "" {
		updatedUser, err := handler.service.SetPartially(ctx.Request.Context(), id, partiallyFieldsToBeSet)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "failed", "error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusOK, updatedUser)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed",
		"message": fmt.Sprintf("And invalid or empty id was specified")})
}

func (handler *Handler) GetAll(ctx *gin.Context) {
	paginationReq := pagination.CreatePaginationRequestFromParams(ctx.Request.URL.Query())
	paginationResponse, err := handler.service.GetAll(ctx.Request.Context(), paginationReq)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, paginationResponse)
}
