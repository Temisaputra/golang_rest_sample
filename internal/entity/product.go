package entity

import "github.com/Temisaputra/warOnk/delivery/presenter"

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

func (p *Products) ToPresenter() *presenter.ProductResponse {
	return &presenter.ProductResponse{
		ProductsId:    p.ID,
		ProductName:   p.ProductName,
		SellingPrice:  p.SellingPrice,
		PurchasePrice: p.PurchasePrice,
		ProductStock:  p.ProductStock,
	}
}
