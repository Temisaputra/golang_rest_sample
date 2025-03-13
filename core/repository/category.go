package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/core/entity"
)

type CategoryRepository interface {
	GetOptionCategory(context.Context) ([]*entity.Category, error)
	GetCategoryByIds(ctx context.Context, ids []int) ([]*entity.Category, error)
}
