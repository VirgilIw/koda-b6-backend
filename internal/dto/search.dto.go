package dto

type SearchProductRequest struct {
	Name              string `form:"name"`
	Category          string `form:"categories"`
	MinPrice          int    `form:"min_price"`
	MaxPrice          int    `form:"max_price"`
	IsFlashSale       bool   `form:"is_flash_sale"`
	IsBuy1Get1        bool   `form:"is_buy1get1"`
	IsBirthdayPackage bool   `form:"is_birthday_package"`
	Cheap             bool   `form:"cheap"`
	Recommended       bool   `form:"recommended"`
	Page              int    `form:"page"`
}

type ProductFilter struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Price             int      `json:"price"`
	FinalPrice        int      `json:"final_price"`
	Categories        []string `json:"categories"`
	Images            []string `json:"images"`
	Rating            *float64 `json:"rating"`
	IsFlashSale       bool     `json:"is_flash_sale"`
	IsBuy1Get1        bool     `json:"is_buy1get1"`
	IsBirthdayPackage bool     `json:"is_birthday_package"`
}
