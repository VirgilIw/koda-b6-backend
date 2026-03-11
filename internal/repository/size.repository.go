package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type SizesRepository struct {
	db *pgxpool.Pool
}

func NewSizesRepository(db *pgxpool.Pool) *SizesRepository {
	return &SizesRepository{
		db: db,
	}
}

func (s *SizesRepository) GetSizes(ctx context.Context) ([]model.Size, error) {
	query := `SELECT id, size_name, created_at, updated_at, deleted_at, additional_price
              FROM sizes;`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Size])
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *SizesRepository) GetSizeById(ctx context.Context, id int) (model.Size, error) {
	query := `SELECT id, size_name, created_at, updated_at, deleted_at, additional_price
              FROM sizes 
			  WHERE ID = $1;`

	rows, err := s.db.Query(ctx, query, id)
	if err != nil {
		return model.Size{}, err
	}
	defer rows.Close()

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Size])
	if err != nil {
		return model.Size{}, err
	}

	return data, nil
}
