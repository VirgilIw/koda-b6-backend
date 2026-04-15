package dto

import "time"

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Error   string  `json:"error,omitempty"`
	Result  []Users `json:"result"`
}

type AuthResponse struct {
	Token string    `json:"token"`
	User  LoginUser `json:"user"`
}

type ResponseOneData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Result  Users  `json:"result"`
}

type ResponseRegister struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Error   string              `json:"error,omitempty"`
	Result  AuthRegisterRequest `json:"result,omitempty"`
}

type ResponseToken struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Error   string       `json:"error,omitempty"`
	Result  AuthResponse `json:"result"`
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
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Sizes       []string `json:"sizes"`
	Stock       int      `json:"stock"`
	Images      *string  `json:"images,omitempty"`
}

type ProductDetailResponse struct {
	Success bool          `json:"success,omitempty"`
	Message string        `json:"message,omitempty"`
	Error   string        `json:"error,omitempty"`
	Result  ProductDetail `json:"result,omitempty"`
}

type ResponseRecommended struct {
	Success bool                         `json:"success,omitempty"`
	Message string                       `json:"message,omitempty"`
	Error   string                       `json:"error,omitempty"`
	Result  []ProductRecommendedResponse `json:"result"`
}

type ProductsResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Total   int         `json:"total"`
	Error   string      `json:"error,omitempty"`
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

type ResponseSizeCreate struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Result  SizeRequest `json:"result"`
}
type ForgotPwdResponse struct {
	Email   string `json:"email"`
	CodeOtp int    `json:"code_otp"`
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

type ResponseReviews struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Error   string    `json:"error,omitempty"`
	Result  []Reviews `json:"result"`
}

type ResponseVariants struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Error   string       `json:"error,omitempty"`
	Result  []AllVariant `json:"result,omitempty"`
}

type ResponseVariant struct {
	Success bool    `json:"success,omitempty"`
	Message string  `json:"message,omitempty"`
	Error   string  `json:"error,omitempty"`
	Result  Variant `json:"result"`
}

type ResponseCreateVariant struct {
	Success bool       `json:"success,omitempty"`
	Message string     `json:"message,omitempty"`
	Error   string     `json:"error,omitempty"`
	Result  AllVariant `json:"result,omitempty"`
}

type ResponseUpdateVariant struct {
	Success bool       `json:"success,omitempty"`
	Message string     `json:"message,omitempty"`
	Error   string     `json:"error,omitempty"`
	Result  AllVariant `json:"result,omitempty"`
}

type ResponseImages struct {
	Success bool       `json:"success,omitempty"`
	Message string     `json:"message,omitempty"`
	Error   string     `json:"error,omitempty"`
	Result  []ImageDto `json:"result,omitempty"`
}

type ResponseImage struct {
	Success bool     `json:"success,omitempty"`
	Message string   `json:"message,omitempty"`
	Error   string   `json:"error,omitempty"`
	Result  ImageDto `json:"result,omitempty"`
}

type ResponseProductFilter struct {
	Success    bool            `json:"success"`
	Message    string          `json:"message"`
	Error      string          `json:"error,omitempty"`
	Page       int             `json:"page"`
	TotalPages int             `json:"total_pages"`
	TotalCount int             `json:"total_count"`
	Result     []ProductFilter `json:"result"`
}

type ResponseCart struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Error   string     `json:"error,omitempty"`
	Result  []CartItem `json:"result"`
}

type TransactionResponse struct {
	ID              int                       `json:"id"`
	TransactionCode string                    `json:"transaction_code"`
	FullName        string                    `json:"full_name"`
	Email           string                    `json:"email"`
	DeliveryMethod  string                    `json:"delivery_method"`
	SubtotalPrice   int                       `json:"subtotal_price"`
	TotalPrice      int                       `json:"total_price"`
	DeliveryFee     int                       `json:"delivery_fee"`
	Tax             int                       `json:"tax"`
	Status          string                    `json:"status"`
	CreatedAt       time.Time                 `json:"created_at"`
	Items           []TransactionItemResponse `json:"items"`
}

type CreateTransactionResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Result  CreateTransactionResult `json:"result"`
}

type GetTransactionResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Result  []TransactionResult `json:"result"`
}

type TransactionResponseHandler struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Result  []TransactionResponse `json:"result"`
}

type GetTransactionDetailResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Result  TransactionDetailResponse `json:"result"`
}
