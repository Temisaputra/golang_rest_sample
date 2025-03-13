package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/core/entity"
)

type ProductRepository interface {
	GetAllProduct(ctx context.Context, pagination *entity.Pagination) ([]*entity.Products, entity.Meta, error)
	GetProductByID(ctx context.Context, id int) (*entity.Products, error)
	CreateProduct(ctx context.Context, product *entity.ProductRequest) error
	UpdateProduct(ctx context.Context, product *entity.ProductRequest, id int) error
	DeleteProduct(ctx context.Context, id int) error
	TransactionalRepository
}
