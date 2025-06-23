package products_repository_mock

import (
	"context"
	"database/sql"

	"github.com/Temisaputra/warOnk/core/dto"
	"github.com/Temisaputra/warOnk/core/repository"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAllProduct(ctx context.Context, pagination *dto.Pagination) ([]*dto.ProductResponse, dto.Meta, error) {
	args := m.Called(ctx, pagination)
	return args.Get(0).([]*dto.ProductResponse), args.Get(1).(dto.Meta), args.Error(2)
}

func (m *MockProductRepository) GetProductByID(ctx context.Context, id int) (*dto.ProductResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*dto.ProductResponse), args.Error(1)
}

func (m *MockProductRepository) CreateProduct(ctx context.Context, product *dto.ProductRequest) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) UpdateProduct(ctx context.Context, product *dto.ProductRequest, id int) error {
	args := m.Called(ctx, product, id)
	return args.Error(0)
}

func (m *MockProductRepository) DeleteProduct(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (repository.TransactionalRepository, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(repository.TransactionalRepository), args.Error(1)
}

func (m *MockProductRepository) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockProductRepository) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockProductRepository) UseTx(tx repository.TransactionalRepository) {
	m.Called(tx)
}
