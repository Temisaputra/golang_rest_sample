package entity

import (
	"github.com/Temisaputra/warOnk/delivery/presenter"
)

type Users struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func (r *Users) TableName() string {
	return "users"
}

func (u *Users) ToPresenter() *presenter.Users {
	return &presenter.Users{
		ID:       u.ID,
		Email:    u.Email,
		Name:     u.Username,
		Password: u.Password,
	}
}
