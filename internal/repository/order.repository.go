package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
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

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Coupon])
	if err != nil {
		return model.Coupon{}, err
	}

	return data, nil
}

func (o *OrderRepository) GetCoupons(ctx context.Context) ([]model.Coupon, error) {
	query := `  SELECT 
		id,
		title,
		description,
		value,
		created_at,
		image
	FROM coupons
`

	rows, err := o.db.Query(ctx, query)
	if err != nil {
		return []model.Coupon{}, err
	}

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Coupon])
	if err != nil {
		return []model.Coupon{}, err
	}

	return data, nil
}

func (o *OrderRepository) CreateCoupon(ctx context.Context, req dto.CouponRequest) (model.Coupon, error) {
	query := `INSERT INTO items (title, description, value, image)
VALUES ($1, $2, $3, $4)
RETURNING id, title, description, value, image`

	rows, err := o.db.Query(ctx, query, req.Title, req.Description, req.Value, req.Image)

	if err != nil {
		return model.Coupon{}, nil
	}

	coupon, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Coupon])

	return coupon, nil
}

func (o *OrderRepository) EditCoupon(ctx context.Context, req dto.CouponRequest) (model.Coupon, error) {
	query := `UPDATE items
SET title = $1,
    description = $2,
    value = $3,
    image = $4
WHERE id = $5
RETURNING id, title, description, value, image;`

	rows, err := o.db.Query(ctx, query, req.Title, req.Description, req.Value, req.Image, req.ID)

	if err != nil {
		return model.Coupon{}, err
	}

	coupon, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Coupon])
	if err != nil {
		return model.Coupon{}, err
	}

	return coupon, nil
}

func (o *OrderRepository) DeleteCoupon(ctx context.Context, id int) (model.Coupon, error) {
	query := `DELETE FROM items
WHERE id = $1
RETURNING id, title, description, value, image, created_at;`

	rows, err := o.db.Query(ctx, query, id)

	if err != nil {
		return model.Coupon{}, nil
	}

	coupon, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Coupon])

	return coupon, nil
}
