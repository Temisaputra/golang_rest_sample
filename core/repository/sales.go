package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/core/entity"
)

type SalesRepository interface {
	CreateSalesHeader(context.Context, *entity.SalesHeaderRequest) (int, error)
	CreateSalesDetail(context.Context, *entity.SalesDetailRequest) error
	TransactionalRepository
}
