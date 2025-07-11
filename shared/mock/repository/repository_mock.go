// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/repository/product.go
//
// Generated by this command:
//
//	mockgen -source=./core/repository/product.go -destination=./shared/mock/repository/repository_mock.go -package repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	dto "github.com/Temisaputra/warOnk/core/dto"
	repository "github.com/Temisaputra/warOnk/core/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
	isgomock struct{}
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockProductRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (repository.TransactionalRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, opts)
	ret0, _ := ret[0].(repository.TransactionalRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockProductRepositoryMockRecorder) BeginTx(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockProductRepository)(nil).BeginTx), ctx, opts)
}

// Commit mocks base method.
func (m *MockProductRepository) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockProductRepositoryMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockProductRepository)(nil).Commit))
}

// CreateProduct mocks base method.
func (m *MockProductRepository) CreateProduct(ctx context.Context, product *dto.ProductRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductRepositoryMockRecorder) CreateProduct(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductRepository)(nil).CreateProduct), ctx, product)
}

// DeleteProduct mocks base method.
func (m *MockProductRepository) DeleteProduct(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductRepositoryMockRecorder) DeleteProduct(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductRepository)(nil).DeleteProduct), ctx, id)
}

// GetAllProduct mocks base method.
func (m *MockProductRepository) GetAllProduct(ctx context.Context, pagination *dto.Pagination) ([]*dto.ProductResponse, dto.Meta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProduct", ctx, pagination)
	ret0, _ := ret[0].([]*dto.ProductResponse)
	ret1, _ := ret[1].(dto.Meta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllProduct indicates an expected call of GetAllProduct.
func (mr *MockProductRepositoryMockRecorder) GetAllProduct(ctx, pagination any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProduct", reflect.TypeOf((*MockProductRepository)(nil).GetAllProduct), ctx, pagination)
}

// GetProductByID mocks base method.
func (m *MockProductRepository) GetProductByID(ctx context.Context, id int) (*dto.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByID", ctx, id)
	ret0, _ := ret[0].(*dto.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByID indicates an expected call of GetProductByID.
func (mr *MockProductRepositoryMockRecorder) GetProductByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByID", reflect.TypeOf((*MockProductRepository)(nil).GetProductByID), ctx, id)
}

// Rollback mocks base method.
func (m *MockProductRepository) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockProductRepositoryMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockProductRepository)(nil).Rollback))
}

// UpdateProduct mocks base method.
func (m *MockProductRepository) UpdateProduct(ctx context.Context, product *dto.ProductRequest, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, product, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductRepositoryMockRecorder) UpdateProduct(ctx, product, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductRepository)(nil).UpdateProduct), ctx, product, id)
}

// UseTx mocks base method.
func (m *MockProductRepository) UseTx(tx repository.TransactionalRepository) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UseTx", tx)
}

// UseTx indicates an expected call of UseTx.
func (mr *MockProductRepositoryMockRecorder) UseTx(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseTx", reflect.TypeOf((*MockProductRepository)(nil).UseTx), tx)
}
