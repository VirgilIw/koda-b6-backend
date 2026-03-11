package dto

type ProductRecommendedResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     int     `json:"price"`
	ImagePath *string `json:"image_path"`
	Rating    float64 `json:"rating"`
	Message   string  `json:"message"`
}
