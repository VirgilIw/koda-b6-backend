package dto

type Product struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	IsBuy1Get1        bool    `json:"is_buyget1"`
	IsFlashSale       bool    `json:"is_flash_sale"`
	IsBirthdayPackage bool    `json:"is_birthday_package"`
}

type ProductsResponse struct {
	Success bool      `json:"success,omitempty"`
	Message string    `json:"message,omitempty"`
	Error   string    `json:"error,omitempty"`
	Result  []Product `json:"result,omitempty"`
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
type CreateProductResponse struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	IsBuy1Get1        bool    `json:"is_buyget1"`
	IsFlashSale       bool    `json:"is_flash_sale"`
	IsBirthdayPackage bool    `json:"is_birthday_package"`
}

// DTO untuk single product response
type SingleProductResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Error   string                `json:"error,omitempty"`
	Result  CreateProductResponse `json:"result"`
}
