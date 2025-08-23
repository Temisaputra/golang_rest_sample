package presenter

type ProductResponse struct {
	ProductsId    int     `json:"products_id"`
	ProductName   string  `json:"product_name"`
	SellingPrice  float64 `json:"selling_price"`
	PurchasePrice float64 `json:"purchase_price"`
	ProductStock  int     `json:"product_stock"`
}

type ProductRequest struct {
	ProductName   string  `json:"product_name"`
	SellingPrice  float64 `json:"selling_price"`
	PurchasePrice float64 `json:"purchase_price"`
	ProductStock  int     `json:"product_stock"`
}
