package module

import (
	"context"
	"fmt"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	"github.com/Temisaputra/warOnk/delivery/repository"
)

type productUsecase struct {
	productRepo     repository.ProductRepository
	transactionRepo repository.TransactionRepository
}

func NewProductUsecase(productRepository repository.ProductRepository, transactionRepository repository.TransactionRepository) *productUsecase {
	return &productUsecase{
		productRepo:     productRepository,
		transactionRepo: transactionRepository,
	}
}

func (u *productUsecase) GetAllProduct(ctx context.Context, pagination *request.Pagination) (produtcs []*presenter.ProductResponse, meta response.Meta, err error) {
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}

	res, meta, err := u.productRepo.GetAllProduct(ctx, pagination)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetAllProduct] Error when getting all product: %w", err)
		return nil, meta, errMsg
	}

	return res, meta, nil
}

func (u *productUsecase) GetProductByID(ctx context.Context, id int) (res *presenter.ProductResponse, err error) {
	res, err = u.productRepo.GetProductByID(ctx, id)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetProductByID] Error when getting product by id: %w", err)
		return nil, errMsg
	}

	return res, nil
}

func (u *productUsecase) CreateProduct(ctx context.Context, params *presenter.ProductRequest) (err error) {
	// Mulai transaksi opsional
	return u.transactionRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Semua repo harus pakai txCtx
		if err := u.productRepo.CreateProduct(txCtx, params); err != nil {
			return err // rollback otomatis
		}

		// if err := u.orderRepo.CreateOrder(txCtx, order); err != nil {
		// 	return err // rollback otomatis
		// }

		return nil // commit otomatis
	})
}

func (u *productUsecase) UpdateProduct(ctx context.Context, params *presenter.ProductRequest, id int) (err error) {
	return u.transactionRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Ambil product dulu, pakai txCtx
		product, err := u.productRepo.GetProductByID(txCtx, id)
		if err != nil {
			return fmt.Errorf("[ProductUsecase-UpdateProduct] Error when getting product by id: %w", err)
		}

		if product == nil {
			return fmt.Errorf("[ProductUsecase-UpdateProduct] Product not found")
		}

		// Update product pakai txCtx
		if err := u.productRepo.UpdateProduct(txCtx, params, id); err != nil {
			return fmt.Errorf("[ProductUsecase-UpdateProduct] Error when updating product: %w", err)
		}

		return nil // commit otomatis kalau tidak ada error
	})
}

func (u *productUsecase) DeleteProduct(ctx context.Context, id int) (err error) {
	return u.transactionRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		// Ambil product dulu, pakai txCtx
		product, err := u.productRepo.GetProductByID(txCtx, id)
		if err != nil {
			return fmt.Errorf("[ProductUsecase-DeleteProduct] Error when getting product by id: %w", err)
		}

		if product == nil {
			return fmt.Errorf("[ProductUsecase-DeleteProduct] Product not found")
		}

		// Delete product pakai txCtx
		if err := u.productRepo.DeleteProduct(txCtx, id); err != nil {
			return fmt.Errorf("[ProductUsecase-DeleteProduct] Error when deleting product: %w", err)
		}

		return nil // commit otomatis kalau tidak ada error
	})
}
