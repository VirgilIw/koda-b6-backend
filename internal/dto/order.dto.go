package dto

import "time"

type Coupon struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Value       string     `json:"value"`
	Image       string     `json:"image"`
	CreatedAt   *time.Time `json:"created_at"`
}
