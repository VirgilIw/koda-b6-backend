package model

type ImageModel struct {
	Id        int    `db:"id"`
	ImagePath string `db:"image_path"`
}
