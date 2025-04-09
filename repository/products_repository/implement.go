package products_repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Temisaputra/warOnk/config"
	"github.com/Temisaputra/warOnk/core/dto"
	irepository "github.com/Temisaputra/warOnk/core/repository"
	"gorm.io/gorm"
)

type repository struct {
	cfg config.Config
	db  *gorm.DB
	tx  *gorm.DB
}

func New(cfg *config.Config, db *gorm.DB) irepository.ProductRepository {
	return &repository{
		cfg: *cfg,
		db:  db,
	}
}

func (r *repository) BeginTx(ctx context.Context, opts *sql.TxOptions) (irepository.TransactionalRepository, error) {
	db := r.db.WithContext(ctx)
	tx := db.Begin(opts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	repo := &repository{
		cfg: r.cfg,
		db:  r.db,
		tx:  tx,
	}

	return repo, nil
}

func (r *repository) Commit() error {
	if r.tx == nil {
		return nil
	}

	err := r.tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Rollback() error {
	if r.tx == nil {
		return nil
	}

	err := r.tx.Rollback().Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UseTx(tx irepository.TransactionalRepository) {
	if txRepo, ok := tx.(*repository); ok {
		r.tx = txRepo.tx
	} else {
		r.tx = nil
	}
}

func (r *repository) GetAllProduct(ctx context.Context, pagination *dto.Pagination) (products []*dto.ProductResponse, meta dto.Meta, err error) {
	db := r.db.WithContext(ctx).Model(&Products{}).Where("deleted_at IS NULL")

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

	var result []Products

	if err = db.Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return nil, meta, err
	}

	meta.Page = int64(pagination.Page)
	meta.PageSize = int64(pagination.PageSize)
	meta.TotalPage = meta.TotalData/int64(pagination.PageSize) + 1

	for _, item := range result {
		products = append(products, item.ToDTO())
	}

	return
}

func (r *repository) GetProductByID(ctx context.Context, id int) (*dto.ProductResponse, error) {
	db := r.db.WithContext(ctx).Model(&Products{}).Where("id = ?", id)
	var result Products
	if err := db.Where("deleted_at IS NULL").First(&result).Error; err != nil {
		return nil, err
	}

	return result.ToDTO(), nil
}

func (r *repository) CreateProduct(ctx context.Context, params *dto.ProductRequest) error {
	conn := r.tx
	if conn == nil {
		conn = r.db.WithContext(ctx)
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	storedData := map[string]interface{}{
		"product_name":   params.ProductName,
		"selling_price":  params.SellingPrice,
		"purchase_price": params.PurchasePrice,
		"product_stock":  params.ProductStock,
		"created_at":     currentTime,
	}
	fmt.Println("data : ", storedData)
	db := conn.WithContext(ctx).Model(&Products{})
	if err := db.Create(&storedData).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateProduct(ctx context.Context, params *dto.ProductRequest, id int) error {
	conn := r.tx
	if conn == nil {
		conn = r.db.WithContext(ctx)
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	updatedData := map[string]interface{}{
		"product_name":   params.ProductName,
		"selling_price":  params.SellingPrice,
		"purchase_price": params.PurchasePrice,
		"product_stock":  params.ProductStock,
		"updated_at":     currentTime,
	}

	db := conn.WithContext(ctx).Model(&Products{})
	if err := db.Where("id = ?", id).Updates(&updatedData).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteProduct(ctx context.Context, id int) error {
	conn := r.tx
	if conn == nil {
		conn = r.db.WithContext(ctx)
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	deletedData := map[string]interface{}{
		"deleted_at": currentTime,
	}

	db := conn.WithContext(ctx).Model(&Products{})
	if err := db.Where("id = ?", id).Updates(&deletedData).Error; err != nil {
		return err
	}

	return nil
}
