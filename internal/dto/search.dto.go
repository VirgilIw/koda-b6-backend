package dto

type SearchProductRequest struct {
	Name              string `form:"name"`
	Category          string `form:"category"`
	MinPrice          int    `form:"min_price"`
	MaxPrice          int    `form:"max_price"`
	IsFlashSale       bool   `form:"is_flash_sale"`
	IsBuy1Get1        bool   `form:"is_buy1get1"`
	IsBirthdayPackage bool   `form:"is_birthday_package"`
	Cheap             bool   `form:"cheap"`
	Page              int    `form:"page"`
	Limit             int    `form:"limit"`
}

type ProductFilter struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Price             int    `json:"price"`
	Category          string `json:"category"`
	IsFlashSale       bool   `json:"is_flash_sale"`
	IsBuy1Get1        bool   `json:"is_buy1get1"`
	IsBirthdayPackage bool   `json:"is_birthday_package"`
}
