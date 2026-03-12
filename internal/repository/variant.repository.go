package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type VariantRepository struct {
	db *pgxpool.Pool
}

func NewVariantRepository(db *pgxpool.Pool) *VariantRepository {
	return &VariantRepository{
		db: db,
	}
}

func (v *VariantRepository) GetVariants(ctx context.Context) ([]model.VariantModel, error) {
	query := `SELECT id, variant_name, additional_price, created_at, updated_at, deleted_at
FROM variants;`

	rows, err := v.db.Query(ctx, query)

	if err != nil {
		return []model.VariantModel{}, err
	}

	datas, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.VariantModel])

	if err != nil {
		return []model.VariantModel{}, err
	}

	return datas, nil
}

func (v *VariantRepository) GetVariantById(ctx context.Context, id int) (model.VariantModel, error) {
	query := `SELECT id, variant_name, additional_price, created_at, updated_at, deleted_at
FROM variants where id = $1`

	rows, err := v.db.Query(ctx, query, id)

	if err != nil {
		return model.VariantModel{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.VariantModel])

	if err != nil {
		return model.VariantModel{}, err
	}

	return data, nil
}
