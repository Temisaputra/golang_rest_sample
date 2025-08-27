package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/internal/domain"
)

type AuthRepository interface {
	Register(ctx context.Context, user *domain.Users) error
	Login(ctx context.Context, email, password string) (*domain.Users, error)
}
