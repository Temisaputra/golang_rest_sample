package sales_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Temisaputra/warOnk/config"
	"github.com/Temisaputra/warOnk/core/entity"
	irepository "github.com/Temisaputra/warOnk/core/repository"
	"gorm.io/gorm"
)

type repository struct {
	cfg config.Config
	db  *gorm.DB
	tx  *gorm.DB
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

func New(cfg *config.Config, db *gorm.DB) irepository.SalesRepository {
	return &repository{
		cfg: *cfg,
		db:  db,
	}
}

func (r *repository) CreateSalesHeader(ctx context.Context, req *entity.SalesHeaderRequest) (int, error) {
	conn := r.tx
	if conn == nil {
		conn = r.db
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userId := 1
	transactionType := 1
	salesHeader := &SalesHeader{
		TransactionDate: currentTime,
		TransactionType: transactionType,
		CustomerName:    req.CustomerName,
		TotalItems:      req.TotalItems,
		TotalAmount:     req.TotalAmount,
		CreatedAt:       &currentTime,
		CreatedBy:       &userId,
		UpdatedAt:       nil,
		UpdatedBy:       nil,
	}

	if err := conn.WithContext(ctx).Model(&SalesHeader{}).Debug().Create(salesHeader).Error; err != nil {
		return 0, err
	}

	return salesHeader.ID, nil
}

func (r *repository) CreateSalesDetail(ctx context.Context, req *entity.SalesDetailRequest) error {
	conn := r.tx
	if conn == nil {
		conn = r.db
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	userId := 1
	salesDetail := &SalesDetail{
		IdSalesHeader: req.IdSalesHeader,
		ProductID:     req.ProductID,
		SalesQuantity: req.SalesQuantity,
		SellingPrice:  req.SalesPrice,
		TotalAmount:   req.TotalAmount,
		CreatedAt:     &currentTime,
		CreatedBy:     &userId,
		UpdatedAt:     nil,
		UpdatedBy:     nil,
	}

	if err := conn.WithContext(ctx).Model(&SalesDetail{}).Debug().Create(salesDetail).Error; err != nil {
		return err
	}

	return nil
}
