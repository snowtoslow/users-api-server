package service

import (
	"context"
	"users-api-server/internal/user"
	"users-api-server/internal/user/model"
	"users-api-server/pkg/pagination"
)

type UserService struct {
	repo user.IUserRepository
}

func New(repo user.IUserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(ctx context.Context, req model.UserReq) (string, error) {
	createdUserID, err := u.repo.CreateUser(ctx, model.UserFromReq(req))
	if err != nil {
		return "", err
	}
	return createdUserID, nil
}

func (u UserService) UpdateUserByID(ctx context.Context, id string, req model.UserReq) (model.UserResponse, error) {
	updatedUserByID, err := u.repo.UpdateUserByID(ctx, id, model.UserFromReq(req))
	if err != nil {
		return model.UserResponse{}, err
	}
	return updatedUserByID.ToResponse(), nil
}

func (u UserService) GetUserByID(ctx context.Context, id string) (model.UserResponse, error) {
	userByID, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return model.UserResponse{}, err
	}
	return userByID.ToResponse(), nil
}

func (u UserService) SetPartially(ctx context.Context, id string, m map[string]interface{}) (model.UserResponse, error) {
	partially, err := u.repo.SetPartially(ctx, id, m)
	if err != nil {
		return model.UserResponse{}, err
	}
	return partially.ToResponse(), nil
}

//http://localhost:8080/admin/all/?page=1&limit=1
//http://localhost:3000/?size=10&page=0&sort=-name
func (u UserService) GetAll(ctx context.Context, request pagination.Request) (model.PaginationResponse, error) {
	users, totalRecords, err := u.repo.GetAll(ctx, request.Page, request.Limit, request.Sort, request.Filter)
	if err != nil {
		return model.PaginationResponse{}, err
	}

	totalPages, nextPage := pagination.CountTotalPagesAndNextPage(totalRecords, request.Limit, request.Page)

	previewPage := make([]model.UserResponse, 0, request.Limit)
	for _, user := range users {
		previewPage = append(previewPage, user.ToResponse())
	}

	return model.PaginationResponse{
		PageInfo: model.PageInfo{
			TotalItems:  totalRecords,
			TotalPage:   totalPages,
			CurrentPage: request.Page,
			NextPage:    nextPage,
		},
		PreviewPage: previewPage,
	}, nil
}
