package dtos

import "time"

type TransactionDto struct {
	ID          int                     `json:"id"`
	TotalAmount int                     `json:"total_amount"`
	CreatedAt   time.Time               `json:"created_at"`
	Details     []TransactionDetailDto  `json:"details"`
}

type TransactionDetailDto struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}
