package handler

import (
	"context"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
)

type userUsecase interface {
	GetAllUser(ctx context.Context, pagination *request.Pagination) (users []*presenter.UserResponse, meta response.Meta, err error)
	GetUserByID(ctx context.Context, id int) (user *presenter.UserResponse, err error)
	CreateUser(ctx context.Context, params *presenter.UserRequest) error
	UpdateUser(ctx context.Context, params *presenter.UserRequest, id int) error
	DeleteUser(ctx context.Context, id int) error
}

type UserHandler struct {
	userUsecase userUsecase
}

func NewUserHandler(userUsecase userUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}
