package dto

import "time"

type Product struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Price             float64   `json:"price"`
	IsFlashSale       bool      `json:"is_flash_sale"`
	IsBuy1Get1        bool      `json:"is_buy1get1"`
	IsBirthdayPackage bool      `json:"is_birthday_package"`
	CreatedAt         time.Time `json:"created_at"`
	ImagePath         *string   `json:"image_path"`
	Rating            float64   `json:"rating"`
}

type ProductDetail struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Price             int      `json:"price"`
	Variants          string   `json:"variants"`
	VariantPrices     []int    `json:"variant_prices"`
	TotalTestimonials int      `json:"total_testimonials"`
	Sizes             []string `json:"sizes"`
	SizePrices        []int    `json:"size_prices"`
}

type UpdateProductRequest struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	IsBuy1Get1        bool    `json:"is_buyget1"`
	IsFlashSale       bool    `json:"is_flash_sale"`
	IsBirthdayPackage bool    `json:"is_birthday_package"`
}

type CreateProductRequest struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	IsBuy1Get1        bool    `json:"is_buyget1"`
	IsFlashSale       bool    `json:"is_flash_sale"`
	IsBirthdayPackage bool    `json:"is_birthday_package"`
}

// DTO untuk single product response
