package model

import "time"

type RecommendedProductModel struct {
	Id             int      `db:"id"`
	Name           string   `db:"name"`
	Description    string   `db:"description"`
	Price          int      `db:"price"`
	ImagePath      *string  `db:"image_path"`
	Rating         *float64 `db:"rating"`
	ReviewMessages *string  `db:"review_messages"`
	TotalReview    int      `db:"total_review"`
}

type Reviews struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Image       string     `db:"image"`
	AuthorTitle string     `db:"author_title"`
	Message     string     `db:"message"`
	Rating      float64    `db:"rating"`
	CreatedAt   *time.Time `db:"created_at"`
	ProductID   int        `db:"product_id"`
}
