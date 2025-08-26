package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/internal/entity"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]presenter.Users, error)
	CreateUser(ctx context.Context, user entity.Users) error
	GetUserByEmail(ctx context.Context, email string) (entity.Users, error)
	UpdateUser(ctx context.Context, user entity.Users) error
	DeleteUser(ctx context.Context, id int) error
}
