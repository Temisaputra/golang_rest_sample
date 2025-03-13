package entity

type Products struct {
	ProductsId         int     `json:"products_id"`
	ProductName        string  `json:"product_name"`
	SellingPrice       float64 `json:"selling_price"`
	PurchasePrice      float64 `json:"purchase_price"`
	Markup             float64 `json:"markup"`
	ProductBarcode     *string `json:"product_barcode"`
	ProductStock       int     `json:"product_stock"`
	ProductImage       *string `json:"product_image"`
	ProductExpiredDate *string `json:"product_expired_date"`
	IdCategoryProduct  int     `json:"id_category_product"`
	IdProductUnit      int     `json:"id_product_unit"`
	MinimumOrderValue  int     `json:"minimum_order_value"`
}

type ProductResponse struct {
	ProductsId         int          `json:"products_id"`
	ProductName        string       `json:"product_name"`
	SellingPrice       float64      `json:"selling_price"`
	PurchasePrice      float64      `json:"purchase_price"`
	Markup             float64      `json:"markup"`
	ProductBarcode     *string      `json:"product_barcode"`
	ProductStock       int          `json:"product_stock"`
	ProductImage       *string      `json:"product_image"`
	ProductExpiredDate *string      `json:"product_expired_date"`
	Category           *Category    `json:"category"`
	Unit               *ProductUnit `json:"unit"`
	MinimumOrderValue  int          `json:"minimum_order_value"`
}

type ProductRequest struct {
	ProductName        string   `json:"product_name"`
	SellingPrice       float64  `json:"selling_price"`
	PurchasePrice      float64  `json:"purchase_price"`
	Markup             float64  `json:"markup"`
	ProductBarcode     string   `json:"product_barcode"`
	ProductStock       int      `json:"product_stock"`
	ProductImage       []string `json:"product_image"`
	ProductExpiredDate string   `json:"product_expired_date"`
	IdCategoryProduct  int      `json:"id_category_product"`
	IdProductUnit      int      `json:"id_product_unit"`
	MinimumOrderValue  int      `json:"minimum_order_value"`
	CreatedAt          *string  `json:"created_at"`
	CreatedBy          *int     `json:"created_by"`
	UpdatedAt          *string  `json:"updated_at"`
	UpdatedBy          *int     `json:"updated_by"`
	DeleteAt           *string  `json:"delete_at"`
	DeleteBy           *int     `json:"delete_by"`
}
