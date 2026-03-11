package model

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
