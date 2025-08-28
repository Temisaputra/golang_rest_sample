package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/internal/domain/entity"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]presenter.UserResponse, error)
	CreateUser(ctx context.Context, user entity.Users) error
	GetUserByEmail(ctx context.Context, email string) (entity.Users, error)
	UpdateUser(ctx context.Context, user entity.Users) error
	DeleteUser(ctx context.Context, id int) error
}
