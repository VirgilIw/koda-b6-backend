package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) GetCouponById(ctx context.Context, id int) (model.Coupon, error) {
	query := `  SELECT 
		id,
		title,
		description,
		value,
		created_at,
		image
	FROM coupons
	WHERE id = $1`

	rows, err := o.db.Query(ctx, query, id)
	if err != nil {
		return model.Coupon{}, err
	}

	defer rows.Close()

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Coupon])
	if err != nil {
		return model.Coupon{}, err
	}

	return data, nil
}
