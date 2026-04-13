package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
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

func (v *VariantRepository) CreateVariant(ctx context.Context, req dto.VariantRequest) (model.VariantModel, error) {
	query := `
	INSERT INTO variants (variant_name, additional_price)
	VALUES ($1, $2)
	RETURNING id, variant_name, additional_price, created_at, updated_at, deleted_at
	`

	rows, err := v.db.Query(ctx, query, req.VariantName, req.AdditionalPrice)
	if err != nil {
		return model.VariantModel{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.VariantModel])
	if err != nil {
		return model.VariantModel{}, err
	}

	return data, nil
}

func (v *VariantRepository) UpdateVariant(ctx context.Context, req dto.VariantRequest, id int) (model.VariantModel, error) {
	query := `
	UPDATE variants
	SET variant_name=$1,
	    additional_price=$2,
	    updated_at=NOW()
	WHERE id=$3
	RETURNING id, variant_name, additional_price, created_at, updated_at, deleted_at
	`

	rows, err := v.db.Query(ctx, query, req.VariantName, req.AdditionalPrice, id)
	if err != nil {
		return model.VariantModel{}, err
	}

	var variant model.VariantModel
	variant, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[model.VariantModel])

	if err != nil {
		return model.VariantModel{}, err
	}

	return variant, nil
}

func (v *VariantRepository) DeleteVariant(ctx context.Context, id int) error {
	query := `DELETE FROM variants
WHERE id=$1`

	cmdTag, err := v.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("variant not found")
	}

	return nil
}

func (v *VariantRepository) GetVariantPrice(ctx context.Context, tx pgx.Tx, variantName string) (int, error) {
	var price int

	err := tx.QueryRow(ctx, `
		SELECT additional_price
		FROM variants
		WHERE variant_name = $1
	`, variantName).Scan(&price)

	if err != nil {
		return 0, err
	}

	return price, nil
}
