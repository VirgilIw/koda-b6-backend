package dto

import "time"

type Size struct {
	ID              int        `json:"id"`
	SizeName        string     `json:"size_name"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	AdditionalPrice int        `json:"additional_price"`
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
	Result  Size   `json:"result,omitempty"`
}
