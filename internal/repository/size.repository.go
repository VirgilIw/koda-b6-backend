package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
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

func (s *SizesRepository) CreateSize(ctx context.Context, req dto.SizeRequest) (model.Size, error) {
	query := `INSERT INTO sizes (size_name, additional_price)
VALUES ($1, $2)
RETURNING id, size_name, created_at, updated_at, deleted_at,  additional_price`
	rows, err := s.db.Query(ctx, query, req.SizeName, req.AdditionalPrice)

	if err != nil {
		return model.Size{}, err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Size])

	if err != nil {
		return model.Size{}, err
	}

	return result, nil
}

func (s *SizesRepository) GetSizes(ctx context.Context) ([]model.Size, error) {
	query := `SELECT id, size_name, created_at, updated_at, deleted_at, additional_price
              FROM sizes;`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

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

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Size])
	if err != nil {
		return model.Size{}, err
	}

	return data, nil
}

func (s *SizesRepository) UpdateSizeById(ctx context.Context, id int, req dto.SizeUpdateRequest) (model.Size, error) {
	query := `UPDATE sizes
SET size_name=$1,
    additional_price=$2,
    updated_at=now()
WHERE id=$3
RETURNING id, size_name, created_at, updated_at, deleted_at, additional_price;`

	rows, err := s.db.Query(ctx, query, req.SizeName, req.AdditionalPrice, id)
	if err != nil {
		return model.Size{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Size])
	if err != nil {
		return model.Size{}, err
	}

	return data, nil
}

func (s *SizesRepository) DeleteSizeById(ctx context.Context, id int) error {
	query := `DELETE FROM sizes WHERE id=$1`

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SizesRepository) GetSizePrice(ctx context.Context, tx pgx.Tx, sizeName string) (int, error) {
	var price int

	err := tx.QueryRow(ctx, `
		SELECT additional_price 
		FROM sizes 
		WHERE size_name = $1
	`, sizeName).Scan(&price)

	if err != nil {
		return 0, err
	}

	return price, nil
}

func (r *ProductRepository) GetProductWithSizes(ctx context.Context, id int) (model.ProductWithSizes, error) {
	query := `
	SELECT 
		p.id,
		p.name,
		p.description,
		p.price,
		p.stock,
		p.created_at,

		COALESCE(
			ARRAY_AGG(s.size_name) FILTER (WHERE s.size_name IS NOT NULL),
			'{}'
		) AS sizes

	FROM products p
	LEFT JOIN product_sizes ps ON ps.product_id = p.id
	LEFT JOIN sizes s ON s.id = ps.size_id

	WHERE p.id = $1
	GROUP BY p.id;
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return model.ProductWithSizes{}, err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ProductWithSizes])
	if err != nil {
		return model.ProductWithSizes{}, err
	}

	return result, nil
}
