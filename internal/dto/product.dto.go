package dto

type Product struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Price             int      `json:"price"`
	Image             string   `json:"image"`
	Rating            float64  `json:"rating"`
	Stock             int      `json:"stock"`
	IsBuy1Get1        bool     `json:"is_buy1_get1"`
	IsFlashSale       bool     `json:"is_flash_sale"`
	IsBirthdayPackage bool     `json:"is_birthday_package"`
	Sizes             []string `json:"sizes"`
}

type SizeResponse struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type VariantResponse struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductDetail struct {
	ID           int               `json:"id"`
	Name         string            `json:"name"`
	Price        int               `json:"price"`
	Description  string            `json:"description"`
	Images       []string          `json:"images"`
	Rating       float64           `json:"rating"`
	TotalReviews int               `json:"total_reviews"`
	Sizes        []SizeResponse    `json:"sizes"`
	Variants     []VariantResponse `json:"variants"`
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
	Name        string  `form:"name"`
	Description string  `form:"description"`
	Price       float64 `form:"price"`
	Stock       int     `form:"stock"`
	Images      *string `form:"images"`
	Sizes       []int64 `json:"sizes[]"`
}
