package user

import (
	"context"
	"users-api-server/internal/user/model"
)

type IUserRepository interface {
	Migrate() error
	CreateUser(context.Context, model.User) (string, error)
	UpdateUserByID(context.Context, string, model.User) (model.User, error)
	GetUserByID(context.Context, string) (model.User, error)
	SetPartially(context.Context, string, map[string]interface{}) (model.User, error)
	GetAll(
		ctx context.Context,
		pageNumber int,
		countInt int,
		sort string,
		filter map[string]interface{},
	) ([]model.User, int64, error)
	//todo: add GetAll with pagination and filters
}
