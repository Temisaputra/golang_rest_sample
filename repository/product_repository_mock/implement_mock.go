package products_repository_mock

import (
	"context"

	"github.com/Temisaputra/warOnk/core/dto"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAllProduct(ctx context.Context, pagination *dto.Pagination) ([]*dto.ProductResponse, dto.Meta, error) {
	args := m.Called(ctx, pagination)
	return args.Get(0).([]*dto.ProductResponse), args.Get(1).(dto.Meta), args.Error(2)
}
