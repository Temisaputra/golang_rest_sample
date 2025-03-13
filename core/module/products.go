package module

import (
	"context"
	"fmt"

	"github.com/Temisaputra/warOnk/core/entity"
	"github.com/Temisaputra/warOnk/core/repository"
)

type productUsecase struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
	unitRepo     repository.UnitRepository
}

func NewProductUsecase(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository, unitRepo repository.UnitRepository) *productUsecase {
	return &productUsecase{
		productRepo:  productRepository,
		categoryRepo: categoryRepository,
		unitRepo:     unitRepo,
	}
}

func (u *productUsecase) GetAllProduct(ctx context.Context, pagination *entity.Pagination) (produtcs []*entity.ProductResponse, meta entity.Meta, err error) {
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}

	products, meta, err := u.productRepo.GetAllProduct(ctx, pagination)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetAllProduct] Error when getting all product: %w", err)
		return nil, meta, errMsg
	}

	categorysId := make([]int, 0, len(products))

	for _, product := range products {
		categorysId = append(categorysId, product.IdCategoryProduct)
	}

	categorys, err := u.categoryRepo.GetCategoryByIds(ctx, categorysId)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetAllProduct] Error when getting category by ids: %w", err)
		return nil, meta, errMsg
	}

	categorysMap := make(map[int]*entity.Category)
	for _, category := range categorys {
		categorysMap[category.CategoryID] = category
	}

	unitsId := make([]int, 0, len(products))

	for _, product := range products {
		unitsId = append(unitsId, product.IdProductUnit)
	}

	units, err := u.unitRepo.GetUnitByIds(ctx, unitsId)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetAllProduct] Error when getting unit by ids: %w", err)
		return nil, meta, errMsg
	}

	unitsMap := make(map[int]*entity.ProductUnit)
	for _, unit := range units {
		unitsMap[unit.UnitID] = unit
	}

	res := make([]*entity.ProductResponse, 0, len(products))
	for _, product := range products {
		category := categorysMap[product.IdCategoryProduct]
		unit := unitsMap[product.IdProductUnit]

		res = append(res, &entity.ProductResponse{
			ProductsId:         product.ProductsId,
			ProductName:        product.ProductName,
			SellingPrice:       product.SellingPrice,
			PurchasePrice:      product.PurchasePrice,
			ProductStock:       product.ProductStock,
			ProductImage:       product.ProductImage,
			ProductExpiredDate: product.ProductExpiredDate,
			Category:           category,
			Unit:               unit,
			MinimumOrderValue:  product.MinimumOrderValue,
		})
	}

	return res, meta, nil
}

func (u *productUsecase) GetProductByID(ctx context.Context, id int) (res *entity.ProductResponse, err error) {
	product, err := u.productRepo.GetProductByID(ctx, id)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetProductByID] Error when getting product by id: %w", err)
		return nil, errMsg
	}

	if product == nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetProductByID] Product not found")
		return nil, errMsg
	}

	category := make([]int, 0, 1)

	category = append(category, product.IdCategoryProduct)
	categorys, err := u.categoryRepo.GetCategoryByIds(ctx, category)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetProductByID] Error when getting category by ids: %w", err)
		return nil, errMsg
	}

	units := make([]int, 0, 1)
	units = append(units, product.IdProductUnit)
	unit, err := u.unitRepo.GetUnitByIds(ctx, units)
	if err != nil {
		errMsg := fmt.Errorf("[ProductUsecase-GetProductByID] Error when getting unit by ids: %w", err)
		return nil, errMsg
	}

	productResponse := &entity.ProductResponse{
		ProductsId:         product.ProductsId,
		ProductName:        product.ProductName,
		SellingPrice:       product.SellingPrice,
		PurchasePrice:      product.PurchasePrice,
		ProductStock:       product.ProductStock,
		ProductImage:       product.ProductImage,
		ProductExpiredDate: product.ProductExpiredDate,
		Category:           categorys[0],
		Unit:               unit[0],
		MinimumOrderValue:  product.MinimumOrderValue,
	}

	return productResponse, nil
}

func (u *productUsecase) CreateProduct(ctx context.Context, params *entity.ProductRequest) (err error) {
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

func (u *productUsecase) UpdateProduct(ctx context.Context, params *entity.ProductRequest, id int) (err error) {
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
