package dto

import "time"

type TransactionItemRequest struct {
	ProductID    int    `json:"product_id" binding:"required"`
	Qty          int    `json:"qty" binding:"required,min=1"`
	Size         string `json:"size" binding:"required"`
	Variant      string `json:"variant" binding:"required"`
	Price        int    `json:"price" validate:"required"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
}

type CreateTransactionRequest struct {
	FullName       string                   `json:"full_name" binding:"required"`
	Email          string                   `json:"email" binding:"required,email"`
	Address        string                   `json:"address" binding:"required"`
	DeliveryMethod string                   `json:"delivery_method" binding:"required"`
	PaymentMethod  string                   `json:"payment_method"`
	Items          []TransactionItemRequest `json:"items" binding:"required,dive"`
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
	DeliveryFee     int        `json:"delivery_fee"`
	Tax             int        `json:"tax"`
	UserId          int        `json:"user_id"`
	TransactionCode string     `json:"transaction_code"`
	Status          string     `json:"status"`
	PaymentMethod   string     `json:"payment_method"`
	CreatedAt       *time.Time `json:"created_at"`
	ProductImage    *string    `json:"image_path"`
}

type CreateTransactionResult struct {
	ID              int    `json:"id"`
	TransactionCode string `json:"transaction_code"`
}

type TransactionItemResponse struct {
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Qty          int    `json:"qty"`
	Size         string `json:"size"`
	Variant      string `json:"variant"`
	Price        int    `json:"price"`
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
