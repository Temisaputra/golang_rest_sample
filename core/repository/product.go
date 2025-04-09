package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/core/dto"
)

type ProductRepository interface {
	GetAllProduct(ctx context.Context, pagination *dto.Pagination) ([]*dto.ProductResponse, dto.Meta, error)
	GetProductByID(ctx context.Context, id int) (*dto.ProductResponse, error)
	CreateProduct(ctx context.Context, product *dto.ProductRequest) error
	UpdateProduct(ctx context.Context, product *dto.ProductRequest, id int) error
	DeleteProduct(ctx context.Context, id int) error
	TransactionalRepository
}
