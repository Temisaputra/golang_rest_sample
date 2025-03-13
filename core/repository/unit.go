package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/core/entity"
)

type UnitRepository interface {
	GetOptionUnit(context.Context) ([]*entity.ProductUnit, error)
	GetUnitByIds(ctx context.Context, ids []int) ([]*entity.ProductUnit, error)
}
