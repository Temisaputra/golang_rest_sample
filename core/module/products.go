package module

import (
	"context"
	"fmt"

	"github.com/Temisaputra/warOnk/core/dto"
	"github.com/Temisaputra/warOnk/core/repository"
)

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) *productUsecase {
	return &productUsecase{
		productRepo: productRepository,
	}
}

func (u *productUsecase) GetAllProduct(ctx context.Context, pagination *dto.Pagination) (produtcs []*dto.ProductResponse, meta dto.Meta, err error) {
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

func (u *productUsecase) GetProductByID(ctx context.Context, id int) (res *dto.ProductResponse, err error) {
	res, err = u.productRepo.GetProductByID(ctx, id)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetProductByID] Error when getting product by id: %w", err)
		return nil, errMsg
	}

	if res == nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetProductByID] Product not found")
		return nil, errMsg
	}

	return res, nil
}

func (u *productUsecase) CreateProduct(ctx context.Context, params *dto.ProductRequest) (err error) {
	tx, err := u.productRepo.BeginTx(ctx, nil)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-CreateProduct] Error when begin transaction: %w", err)
		return errMsg
	}

	u.productRepo.UseTx(tx)

	err = u.productRepo.CreateProduct(ctx, params)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Errorf("[ProductUsecase-CreateProduct] Error when creating product: %w", err)
		return errMsg
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Errorf("[ProductUsecase-CreateProduct] Error when commit transaction: %w", err)
		return errMsg
	}

	return nil
}

func (u *productUsecase) UpdateProduct(ctx context.Context, params *dto.ProductRequest, id int) (err error) {
	tx, err := u.productRepo.BeginTx(ctx, nil)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-UpdateProduct] Error when begin transaction: %w", err)
		return errMsg
	}

	product, err := u.productRepo.GetProductByID(ctx, id)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-UpdateProduct] Error when getting product by id: %w", err)
		return errMsg
	}

	if product == nil {
		errMsg := fmt.Errorf("[ProductUsecase-UpdateProduct] Product not found")
		return errMsg
	}

	u.productRepo.UseTx(tx)

	err = u.productRepo.UpdateProduct(ctx, params, id)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Errorf("[ProductUsecase-UpdateProduct] Error when updating product: %w", err)
		return errMsg
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Errorf("[ProductUsecase-UpdateProduct] Error when commit transaction: %w", err)
		return errMsg
	}

	return nil
}

func (u *productUsecase) DeleteProduct(ctx context.Context, id int) (err error) {
	tx, err := u.productRepo.BeginTx(ctx, nil)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-DeleteProduct] Error when begin transaction: %w", err)
		return errMsg
	}

	product, err := u.productRepo.GetProductByID(ctx, id)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-UpdateProduct] Error when getting product by id: %w", err)
		return errMsg
	}

	if product == nil {
		errMsg := fmt.Errorf("[ProductUsecase-UpdateProduct] Product not found")
		return errMsg
	}

	u.productRepo.UseTx(tx)

	err = u.productRepo.DeleteProduct(ctx, id)
	if err != nil {
		tx.Rollback()
		errMsg := fmt.Errorf("[ProductUsecase-DeleteProduct] Error when deleting product: %w", err)
		return errMsg
	}

	err = tx.Commit()
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-DeleteProduct] Error when commit transaction: %w", err)
		return errMsg
	}

	return nil
}
