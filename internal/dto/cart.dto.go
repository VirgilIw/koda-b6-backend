package dto

import "time"

type AddToCartRequest struct {
	UserID    int `json:"user_id" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
	Qty       int `json:"qty" binding:"required"`
	SizeID    int `json:"size_id"`
	VariantID int `json:"variant_id"`
}

type CartItem struct {
	ID           int        `json:"id"`
	UserID       int        `json:"user_id"`
	ProductID    int        `json:"product_id"`
	ProductName  string     `json:"product_name"`
	ProductImage string     `json:"product_image"`
	Price        int        `json:"price"`
	Qty          int        `json:"qty"`
	SizeID       int        `json:"size_id"`
	VariantID    int        `json:"variant_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
