package dto

type TransactionItemRequest struct {
	ProductID int    `json:"product_id" validate:"required"`
	Qty       int    `json:"qty" validate:"required,min=1"`
	Size      string `json:"size" validate:"required"`
	Variant   string `json:"variant" validate:"required"`
	Price     int    `json:"price" validate:"required"`
}

type CreateTransactionRequest struct {
	FullName       string                   `json:"full_name" validate:"required"`
	Email          string                   `json:"email" validate:"required,email"`
	Address        string                   `json:"address" validate:"required"`
	DeliveryMethod string                   `json:"delivery_method" validate:"required"`
	SubtotalPrice  int                      `json:"subtotal_price" validate:"required"`
	TotalPrice     int                      `json:"total_price" validate:"required"`
	DeliveryFee    int                      `json:"delivery_fee" validate:"required"`
	Tax            int                      `json:"tax" validate:"required"`
	PaymentMethod  string                   `json:"payment_method"`
	Items          []TransactionItemRequest `json:"items" validate:"required,dive"`
}

type UpdateTransactionStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
type TransactionResult struct {
	TransactionCode string `json:"transaction_code"`
}
