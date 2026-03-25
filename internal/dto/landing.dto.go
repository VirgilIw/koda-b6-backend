package dto

import "time"

type ProductRecommendedResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	ImagePath   *string `json:"image_path"`
	Rating      float64 `json:"rating"`
	Message     string  `json:"message"`
}

type Reviews struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Image       string     `db:"image"`
	AuthorTitle string     `db:"author_title"`
	Message     string     `db:"message"`
	Rating      float64    `db:"rating"`
	CreatedAt   *time.Time `db:"created_at"`
}
