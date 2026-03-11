package dto

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Error   string  `json:"error,omitempty"`
	Result  []Users `json:"result"`
}

type ResponseOneData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  Users  `json:"result"`
}

type ResponseToken struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Token   string `json:"token,omitempty"`
}

type ResponseError struct {
	Status  string `json:"status" example:"Error"`
	Message string `json:"message" example:"Failed get data"`
	Error   string `json:"errors,omitempty" example:"failed get data"`
}

type SingleProductResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Error   string                `json:"error,omitempty"`
	Result  CreateProductResponse `json:"result"`
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

type ProductDetailResponse struct {
	Success bool          `json:"success,omitempty"`
	Message string        `json:"message,omitempty"`
	Error   string        `json:"error,omitempty"`
	Result  ProductDetail `json:"result,omitempty"`
}

type ProductsResponse struct {
	Success bool      `json:"success,omitempty"`
	Message string    `json:"message,omitempty"`
	Error   string    `json:"error,omitempty"`
	Result  []Product `json:"result,omitempty"`
}

type ResponseCoupon struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Result  Coupon `json:"result"`
}

type ResponseAllCoupon struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Error   string   `json:"error,omitempty"`
	Result  []Coupon `json:"result"`
}

type ResponseOneCoupon struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Result  Coupon `json:"result"`
}

type ResponseCategories struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Error   string     `json:"error,omitempty"`
	Result  []Category `json:"result"`
}

type ResponseCategory struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Error   string   `json:"error,omitempty"`
	Result  Category `json:"result"`
}

type ResponseSizes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Result  []Size `json:"result,omitempty"`
}

type ResponseSize struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Result  Size   `json:"result"`
}

type ForgotPwdResponse struct {
	Email   string
	CodeOtp int
}

type ResponseForgotPwd struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Error   string            `json:"error,omitempty"`
	Result  ForgotPwdResponse `json:"result"`
}

type ResponseResetPwd struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
