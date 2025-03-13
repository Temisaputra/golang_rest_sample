package unit_repository

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

func New(cfg *config.Config, db *gorm.DB) irepository.UnitRepository {
	return &repository{
		cfg: *cfg,
		db:  db,
	}
}

func (r *repository) GetOptionUnit(ctx context.Context) (res []*entity.ProductUnit, err error) {
	db := r.db.WithContext(ctx).Model(&ProductUnit{}).Order("id ASC")

	var result []ProductUnit
	if err = db.Find(&result).Error; err != nil {
		return nil, err
	}

	for _, item := range result {
		res = append(res, item.ToEntity())
	}

	return res, nil
}

func (r *repository) GetUnitByIds(ctx context.Context, ids []int) (res []*entity.ProductUnit, err error) {
	db := r.db.WithContext(ctx).Model(&ProductUnit{}).Where("id IN (?)", ids)

	var result []ProductUnit
	if err = db.Find(&result).Error; err != nil {
		return nil, err
	}

	for _, item := range result {
		res = append(res, item.ToEntity())
	}

	return res, nil
}
