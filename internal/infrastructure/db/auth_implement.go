package db

import (
	"context"

	irepository "github.com/Temisaputra/warOnk/delivery/repository"
	"github.com/Temisaputra/warOnk/internal/domain"
	"gorm.io/gorm"
)

type AuthRepository struct {
	*TransactionRepository
}

func NewAuthRepo(db *gorm.DB) irepository.AuthRepository {
	return &AuthRepository{
		TransactionRepository: NewTransactionRepo(db),
	}
}

func (r *AuthRepository) Register(ctx context.Context, user *domain.Users) error {
	if err := r.Conn(ctx).WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) Login(ctx context.Context, email, password string) (*domain.Users, error) {
	var user domain.Users
	if err := r.Conn(ctx).WithContext(ctx).Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
