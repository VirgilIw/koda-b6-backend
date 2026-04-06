package dto

import "time"

type TransactionItemRequest struct {
	ProductID int    `json:"product_id" validate:"required"`
	Qty       int    `json:"qty" validate:"required,min=1"`
	Size      string `json:"size" validate:"required"`
	Variant   string `json:"variant" validate:"required"`
	Price     int    `json:"price" validate:"required"`
}

type CreateTransactionRequest struct {
	FullName       string                   `json:"full_name" validate:"required"`
	Email          string                   `json:"email" validate:"required,email"`
	Address        string                   `json:"address" validate:"required"`
	DeliveryMethod string                   `json:"delivery_method" validate:"required"`
	SubtotalPrice  int                      `json:"subtotal_price" validate:"required"`
	TotalPrice     int                      `json:"total_price" validate:"required"`
	DeliveryFee    int                      `json:"delivery_fee" validate:"required"`
	Tax            int                      `json:"tax" validate:"required"`
	PaymentMethod  string                   `json:"payment_method"`
	Items          []TransactionItemRequest `json:"items" validate:"required,dive"`
}

type UpdateTransactionStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
type TransactionResult struct {
	ID              int        `json:"id"`
	DeliveryMethod  string     `json:"delivery_method"`
	FullName        string     `json:"full_name"`
	Email           string     `json:"email"`
	Address         string     `json:"address"`
	SubtotalPrice   int        `json:"subtotal_price"`
	TotalPrice      int        `json:"total_price"`
	DeliveryFee     string     `json:"delivery_fee"`
	Tax             int        `json:"tax"`
	TransactionCode string     `json:"transaction_code"`
	Status          string     `json:"status"`
	PaymentMethod   int        `json:"payment_method"`
	CreatedAt       *time.Time `json:"created_at"`
}

type CreateTransactionResult struct {
	ID              int    `json:"id"`
	TransactionCode string `json:"transaction_code"`
}

type TransactionItemResponse struct {
	ProductID int    `json:"product_id"`
	Qty       int    `json:"qty"`
	Size      string `json:"size"`
	Variant   string `json:"variant"`
	Price     int    `json:"price"`
}

type TransactionDetailResponse struct {
	ID              int                       `json:"id"`
	TransactionCode string                    `json:"transaction_code"`
	FullName        string                    `json:"full_name"`
	Email           string                    `json:"email"`
	Address         string                    `json:"address"`
	DeliveryMethod  string                    `json:"delivery_method"`
	SubtotalPrice   int                       `json:"subtotal_price"`
	TotalPrice      int                       `json:"total_price"`
	DeliveryFee     int                       `json:"delivery_fee"`
	Tax             int                       `json:"tax"`
	PaymentMethod   string                    `json:"payment_method"`
	Status          string                    `json:"status"`
	CreatedAt       *time.Time                `json:"created_at"`
	Items           []TransactionItemResponse `json:"items"`
}
