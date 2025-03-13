package sales_repository

type SalesHeader struct {
	ID              int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	TransactionDate string  `json:"transaction_date" gorm:"column:transaction_date"`
	TransactionType int     `json:"transaction_type" gorm:"column:transaction_type"`
	CustomerName    *string `json:"customer_name" gorm:"column:customer_name"`
	TotalItems      int     `json:"total_items" gorm:"column:total_items"`
	TotalAmount     float64 `json:"total_amount" gorm:"column:total_amount"`
	CreatedAt       *string `json:"created_at" gorm:"column:created_at"`
	CreatedBy       *int    `json:"created_by" gorm:"column:created_by"`
	UpdatedAt       *string `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy       *int    `json:"updated_by" gorm:"column:updated_by"`
}

func (s SalesHeader) TableName() string {
	return "sales_header"
}

type SalesDetail struct {
	ID            int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	IdSalesHeader int     `json:"id_sales_header" gorm:"column:id_sales_header"`
	ProductID     int     `json:"product_id" gorm:"column:product_id"`
	SalesQuantity int     `json:"sales_quantity" gorm:"column:sales_quantity"`
	SellingPrice  float64 `json:"selling_price" gorm:"column:selling_price"`
	TotalAmount   float64 `json:"total_amount" gorm:"column:total_amount"`
	CreatedAt     *string `json:"created_at" gorm:"column:created_at"`
	CreatedBy     *int    `json:"created_by" gorm:"column:created_by"`
	UpdatedAt     *string `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy     *int    `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt     *string `json:"deleted_at" gorm:"column:deleted_at"`
	DeleteBy      *int    `json:"deleted_by" gorm:"column:deleted_by"`
}

func (s SalesDetail) TableName() string {
	return "sales_details"
}
