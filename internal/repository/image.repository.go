package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ImageRepository struct {
	db *pgxpool.Pool
}

func NewImageRepository(db *pgxpool.Pool) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}
