package model

import "time"

type ProductModel struct {
	ID                int       `db:"id"`
	Name              string    `db:"name"`
	Description       string    `db:"description"`
	Price             int       `db:"price"`
	Image             string    `db:"image"`
	Rating            *float64  `db:"rating"`
	IsFlashSale       bool      `db:"is_flash_sale"`
	IsBuy1Get1        bool      `db:"is_buy1get1"`
	IsBirthdayPackage bool      `db:"is_birthday_package"`
	CreatedAt         time.Time `db:"created_at"`
}

type ProductDetail struct {
	ID                int      `db:"id"`
	Name              string   `db:"name"`
	Price             int      `db:"price"`
	Description       string   `db:"description"`
	Images            []string `db:"images"`
	Variants          []string `db:"variants"`
	VariantPrices     []int    `db:"variant_prices"`
	TotalTestimonials int      `db:"total_testimonials"`
	Sizes             []string `db:"sizes"`
	SizePrices        []int    `db:"size_prices"`
}
