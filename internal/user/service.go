package user

import (
	"context"
	"users-api-server/internal/user/model"
	"users-api-server/pkg/pagination"
)

type IUserService interface {
	CreateUser(context.Context, model.UserReq) (string, error)
	UpdateUserByID(context.Context, string, model.UserReq) (model.UserResponse, error)
	GetUserByID(context.Context, string) (model.UserResponse, error)
	SetPartially(context.Context, string, map[string]interface{}) (model.UserResponse, error)
	GetAll(ctx context.Context, request pagination.Request) (model.PaginationResponse, error)
}
