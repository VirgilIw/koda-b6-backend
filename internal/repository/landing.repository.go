package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type LandingRepository struct {
	db *pgxpool.Pool
}

func NewLandingRepository(db *pgxpool.Pool) *LandingRepository {
	return &LandingRepository{
		db: db,
	}
}

func (r *LandingRepository) GetReviews(ctx context.Context) ([]model.Reviews, error) {
	query := `SELECT id, "name", image, author_title, message, rating, created_at, product_id
FROM testimonials order by id;`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return []model.Reviews{}, err
	}

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Reviews])

	if err != nil {
		return []model.Reviews{}, err
	}

	return data, nil
}
