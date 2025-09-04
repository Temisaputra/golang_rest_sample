package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/internal/delivery/presenter"
	"github.com/Temisaputra/warOnk/internal/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/internal/delivery/presenter/response"
)

type ProductRepository interface {
	GetAllProduct(ctx context.Context, pagination *request.Pagination) ([]*presenter.ProductResponse, response.Meta, error)
	GetProductByID(ctx context.Context, id int) (*presenter.ProductResponse, error)
	CreateProduct(ctx context.Context, product *presenter.ProductRequest) error
	UpdateProduct(ctx context.Context, product *presenter.ProductRequest, id int) error
	DeleteProduct(ctx context.Context, id int) error
}
