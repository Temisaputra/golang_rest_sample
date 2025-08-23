package db

import (
	"context"
	"strings"
	"time"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	irepository "github.com/Temisaputra/warOnk/delivery/repository"
	"github.com/Temisaputra/warOnk/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	*TransactionRepository
}

func NewProductRepo(db *gorm.DB) irepository.ProductRepository {
	return &ProductRepository{
		TransactionRepository: NewTransactionRepo(db),
	}
}

func (r *ProductRepository) GetAllProduct(ctx context.Context, pagination *request.Pagination) (products []*presenter.ProductResponse, meta response.Meta, err error) {
	db := r.Conn(ctx).WithContext(ctx).Model(&entity.Products{}).Where("deleted_at IS NULL")

	if pagination.Keyword != "" {
		keywordStr := "%" + pagination.Keyword + "%"
		db = db.Where("id ILIKE ? OR product_name ILIKE ? ", keywordStr, keywordStr)
	}

	if pagination.OrderBy != "" || pagination.OrderType != "" {
		splitOrder := strings.Split(pagination.OrderBy, "|")
		if len(splitOrder) > 1 {
			db = db.Order(splitOrder[0] + " " + splitOrder[1] + " " + pagination.OrderType)
		} else {
			db = db.Order(pagination.OrderBy + " " + pagination.OrderType)
		}
	} else {
		db = db.Order("updated_at DESC NULLS LAST")
	}

	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	if err = db.Count(&meta.TotalData).Error; err != nil {
		return nil, meta, err
	}

	var result []entity.Products

	if err = db.Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return nil, meta, err
	}

	meta.Page = int64(pagination.Page)
	meta.PageSize = int64(pagination.PageSize)
	meta.TotalPage = meta.TotalData/int64(pagination.PageSize) + 1

	for _, item := range result {
		products = append(products, item.ToPresenter())
	}

	return
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*presenter.ProductResponse, error) {
	db := r.Conn(ctx).WithContext(ctx).Model(&entity.Products{}).Where("id = ?", id)
	var result entity.Products
	if err := db.Where("deleted_at IS NULL").First(&result).Error; err != nil {
		return nil, err
	}

	return result.ToPresenter(), nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, params *presenter.ProductRequest) error {
	conn := r.Conn(ctx)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	storedData := map[string]interface{}{
		"product_name":   params.ProductName,
		"selling_price":  params.SellingPrice,
		"purchase_price": params.PurchasePrice,
		"product_stock":  params.ProductStock,
		"created_at":     currentTime,
	}
	db := conn.WithContext(ctx).Model(&entity.Products{})
	if err := db.Create(&storedData).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, params *presenter.ProductRequest, id int) error {
	conn := r.Conn(ctx)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	updatedData := map[string]interface{}{
		"product_name":   params.ProductName,
		"selling_price":  params.SellingPrice,
		"purchase_price": params.PurchasePrice,
		"product_stock":  params.ProductStock,
		"updated_at":     currentTime,
	}

	db := conn.WithContext(ctx).Model(&entity.Products{})
	if err := db.Where("id = ?", id).Updates(&updatedData).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	conn := r.Conn(ctx).WithContext(ctx)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	deletedData := map[string]interface{}{
		"deleted_at": currentTime,
	}

	db := conn.WithContext(ctx).Model(&entity.Products{})
	if err := db.Where("id = ?", id).Updates(&deletedData).Error; err != nil {
		return err
	}

	return nil
}
