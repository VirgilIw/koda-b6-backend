package dto

type ImageDto struct {
	Id        int    `json:"id"`
	ImagePath string `json:"image_path"`
}

type ImageRequest struct {
	ImagePath string `json:"image_path"`
}
