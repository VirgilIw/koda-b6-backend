package model

import "time"

type CartItem struct {
	ID           int        `db:"id"`
	UserID       int        `db:"user_id"`
	ProductID    int        `db:"product_id"`
	ProductName  string     `db:"name"`
	ProductImage string     `db:"image_path"`
	Price        int        `db:"price"`
	Qty          int        `db:"qty"`
	SizeID       int        `db:"size_id"`
	VariantID    int        `db:"variant_id"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
