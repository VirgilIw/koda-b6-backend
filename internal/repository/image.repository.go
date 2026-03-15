package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type ImageRepository struct {
	db *pgxpool.Pool
}

func NewImageRepository(db *pgxpool.Pool) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

func (i *ImageRepository) CreateImage(ctx context.Context, req dto.ImageRequest) (model.ImageModel, error) {
	query := `INSERT INTO images (image_path) VALUES ($1) RETURNING id, image_path`

	rows, err := i.db.Query(ctx, query, req.ImagePath)

	if err != nil {
		return model.ImageModel{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ImageModel])

	return data, nil
}

func (i *ImageRepository) GetImages(ctx context.Context) ([]model.ImageModel, error) {
	query := `select id, image_path from images order by id asc`

	rows, err := i.db.Query(ctx, query)

	if err != nil {
		return []model.ImageModel{}, err
	}

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.ImageModel])

	if err != nil {
		return []model.ImageModel{}, err
	}

	return data, nil
}

func (i *ImageRepository) GetImageById(ctx context.Context, id int) (model.ImageModel, error) {
	query := `select id, image_path from images where id = $1`

	rows, err := i.db.Query(ctx, query, id)

	if err != nil {
		return model.ImageModel{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ImageModel])

	if err != nil {
		return model.ImageModel{}, err
	}

	return data, nil
}

func (i *ImageRepository) UpdateImage(ctx context.Context, req dto.ImageRequest, id int) (model.ImageModel, error) {
	query := `update images set image_path = $1 where id = $2 returning id ,image_path`

	rows, err := i.db.Query(ctx, query, req.ImagePath, id)

	if err != nil {
		return model.ImageModel{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ImageModel])

	if err != nil {
		return model.ImageModel{}, err
	}

	return data, nil
}

func (i *ImageRepository) DeleteImageById(ctx context.Context, id int) error {
	query := `DELETE FROM images WHERE id = $1`

	result, err := i.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rows := result.RowsAffected()

	if rows == 0 {
		return errors.New("id not found")
	}

	return nil
}
