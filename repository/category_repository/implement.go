package category_repository

import (
	"context"

	"github.com/Temisaputra/warOnk/config"
	"github.com/Temisaputra/warOnk/core/entity"
	irepository "github.com/Temisaputra/warOnk/core/repository"
	"gorm.io/gorm"
)

type repository struct {
	cfg config.Config
	db  *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) irepository.CategoryRepository {
	return &repository{
		cfg: *cfg,
		db:  db,
	}
}

func (r *repository) GetOptionCategory(ctx context.Context) (res []*entity.Category, err error) {
	db := r.db.WithContext(ctx).Model(&Category{}).Order("id ASC")

	var result []Category
	if err = db.Find(&result).Error; err != nil {
		return nil, err
	}

	for _, item := range result {
		res = append(res, item.ToEntity())
	}

	return res, nil
}

func (r *repository) GetCategoryByIds(ctx context.Context, ids []int) (res []*entity.Category, err error) {
	db := r.db.WithContext(ctx).Model(&Category{}).Where("id IN (?)", ids)

	var result []Category
	if err = db.Find(&result).Error; err != nil {
		return nil, err
	}

	for _, item := range result {
		res = append(res, item.ToEntity())
	}

	return res, nil
}
