package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type UserRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewUserRepository(db *pgxpool.Pool, rdb *redis.Client) *UserRepository {
	return &UserRepository{
		db:  db,
		rdb: rdb,
	}
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]model.UserModel, error) {
	cachedKey := "users:all"
	valueCache, err := u.rdb.Get(ctx, cachedKey).Result()

	if err == redis.Nil {
		query := `
SELECT 
    id,
    fullname,
    email,
    password,
    phone,
    address,
    picture,
    role,
    created_at,
    updated_at,
    deleted_at,
   	lastlogin_at
FROM users;
	`
		rows, err := u.db.Query(ctx, query)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.UserModel])

		val, err := json.Marshal(users)
		u.rdb.Set(ctx, cachedKey, string(val), time.Minute*15)

		return users, nil
	} else if err != nil {
		return nil, err
	} else {
		users := []model.UserModel{}
		if err := json.Unmarshal([]byte(valueCache), &users); err != nil {
			return nil, err
		}
		return users, nil
	}
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (model.UserModel, error) {
	query := `
SELECT 
    id,
    fullname,
    email,
    password,
    COALESCE(phone, '') AS phone,
    COALESCE(address, '') AS address,
    COALESCE(picture, '') AS picture,
    COALESCE(role, '') AS role,
    COALESCE(created_at, NOW()) AS created_at,
    COALESCE(updated_at, NOW()) AS updated_at,
    deleted_at,
    COALESCE(lastlogin_at, NOW()) AS lastlogin_at
FROM users
WHERE email = $1;`

	var user model.UserModel

	err := u.db.QueryRow(ctx, query, email).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Address,
		&user.Picture,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.LastLoginAt,
	)

	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}
