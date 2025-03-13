package entity

import "time"

type SalesRequest struct {
	SalesHeader *SalesHeaderRequest   `json:"sales_header"`
	SalesDetail *[]SalesDetailRequest `json:"sales_detail"`
}

type SalesHeaderRequest struct {
	TransactionDate string     `json:"transaction_date"`
	TransactionType int        `json:"transaction_type"`
	CustomerName    *string    `json:"customer_name"`
	TotalItems      int        `json:"total_items"`
	TotalAmount     float64    `json:"total_amount"`
	CreatedBy       *int       `json:"created_by"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedBy       *int       `json:"updated_by"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type SalesDetailRequest struct {
	IdSalesHeader int     `json:"id_sales_header"`
	ProductID     int     `json:"product_id"`
	SalesQuantity int     `json:"sales_quantity"`
	SalesPrice    float64 `json:"sales_price"`
	TotalAmount   float64 `json:"total_amount"`
	CreatedBy     *int    `json:"created_by"`
	CreatedAt     *string `json:"created_at"`
	UpdatedBy     *int    `json:"updated_by"`
	UpdatedAt     *string `json:"updated_at"`
	DeletedBy     *int    `json:"deleted_by"`
	DeletedAt     *string `json:"deleted_at"`
}
