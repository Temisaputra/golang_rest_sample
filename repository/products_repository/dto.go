package products_repository

import "github.com/Temisaputra/warOnk/core/entity"

type Products struct {
	ID                 int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProductName        string  `json:"product_name" gorm:"column:product_name"`
	SellingPrice       float64 `json:"selling_price" gorm:"column:selling_price"`
	PurchasePrice      float64 `json:"purchase_price" gorm:"column:purchase_price"`
	Markup             float64 `json:"markup" gorm:"column:markup"`
	ProductExpiredDate *string `json:"product_expired_date" gorm:"column:product_expired_date"`
	ProductBarcode     *string `json:"product_barcode" gorm:"column:product_barcode"`
	ProductStock       int     `json:"product_stock" gorm:"column:product_stock"`
	ProductImage       *string `json:"product_image" gorm:"column:product_image"`
	IDCategoryProduct  int     `json:"id_category_product" gorm:"column:id_category_product"`
	IDProductUnit      int     `json:"id_product_unit" gorm:"column:id_product_unit"`
	MinimumOrderValue  int     `json:"minimum_order_value" gorm:"column:minimum_order_value"`
	CreatedAt          *string `json:"created_at" gorm:"column:created_at"`
	CreatedBy          *string `json:"created_by" gorm:"column:created_by"`
	UpdatedAt          *string `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy          *string `json:"updated_by" gorm:"column:updated_by"`
	DeleteAt           *string `json:"delete_at" gorm:"column:delete_at"`
	DeleteBy           *string `json:"delete_by" gorm:"column:delete_by"`
}

func (Products) TableName() string {
	return "product"
}

func (p *Products) ToEntity() *entity.Products {
	return &entity.Products{
		ProductsId:         p.ID,
		ProductName:        p.ProductName,
		SellingPrice:       p.SellingPrice,
		PurchasePrice:      p.PurchasePrice,
		Markup:             p.Markup,
		ProductExpiredDate: p.ProductExpiredDate,
		ProductBarcode:     p.ProductBarcode,
		ProductStock:       p.ProductStock,
		ProductImage:       p.ProductImage,
		IdCategoryProduct:  p.IDCategoryProduct,
		IdProductUnit:      p.IDProductUnit,
		MinimumOrderValue:  p.MinimumOrderValue,
	}
}
