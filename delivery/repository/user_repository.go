package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/internal/domain"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]presenter.UserResponse, error)
	CreateUser(ctx context.Context, user domain.Users) error
	GetUserByEmail(ctx context.Context, email string) (domain.Users, error)
	UpdateUser(ctx context.Context, user domain.Users) error
	DeleteUser(ctx context.Context, id int) error
}
