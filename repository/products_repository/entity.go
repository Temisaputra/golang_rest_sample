package products_repository

import (
	"github.com/Temisaputra/warOnk/core/dto"
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Products struct {
	ID            int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProductName   string  `json:"product_name" gorm:"column:product_name"`
	SellingPrice  float64 `json:"selling_price" gorm:"column:selling_price"`
	PurchasePrice float64 `json:"purchase_price" gorm:"column:purchase_price"`
	ProductStock  int     `json:"product_stock" gorm:"column:product_stock"`
	CreatedAt     string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     string  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     string  `json:"deleted_at" gorm:"column:deleted_at"`
}

func (r *Products) TableName() string {
	return "products"
}

func (p *Products) ToDTO() *dto.ProductResponse {
	return &dto.ProductResponse{
		ProductsId:    p.ID,
		ProductName:   p.ProductName,
		SellingPrice:  p.SellingPrice,
		PurchasePrice: p.PurchasePrice,
		ProductStock:  p.ProductStock,
	}
}

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20240611_create_products",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Products{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("products")
			},
		},
	})

	return m.Migrate()
}
