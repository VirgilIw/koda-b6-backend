package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
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

func (u *UserRepository) invalidateUserCache(ctx context.Context) {
	u.rdb.Del(ctx, "users:all")
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]model.UserModel, error) {
	cachedKey := "users:all"

	valueCache, err := u.rdb.Get(ctx, cachedKey).Result()

	if err == redis.Nil {
		query := `
SELECT 
    id,
    COALESCE(NULLIF(fullname, ''), 'No Name') AS fullname,
    email,
    password,
    phone,
    address,
    picture,
    COALESCE(NULLIF(role, ''), 'No Role') AS role,
    created_at,
    updated_at,
    deleted_at,
    lastlogin_at
FROM users
ORDER BY id ASC;
		`

		rows, err := u.db.Query(ctx, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.UserModel])
		if err != nil {
			return nil, err
		}

		//cache hanya kalau ada data
		if len(users) > 0 {
			val, err := json.Marshal(users)
			if err == nil {
				u.rdb.Set(ctx, cachedKey, val, time.Minute*15)
			}
		}

		return users, nil

	} else if err != nil {
		return nil, err

	} else {
		users := []model.UserModel{}
		if err := json.Unmarshal([]byte(valueCache), &users); err != nil {
			u.rdb.Del(ctx, cachedKey)
			return u.GetUsers(ctx)
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
    phone,
    address,
    picture,
    role,
    created_at,
    updated_at,
    deleted_at,
    lastlogin_at
FROM users
WHERE email = $1;
`

	rows, err := u.db.Query(ctx, query, email)
	if err != nil {
		return model.UserModel{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.UserModel])
	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (u *UserRepository) GetById(ctx context.Context, id int) (model.UserModel, error) {
	query := `
SELECT 
    id, fullname, email, password, phone, address, picture, role, 
    created_at, updated_at, deleted_at, lastlogin_at
FROM users
WHERE id = $1;
`

	var user model.UserModel
	err := u.db.QueryRow(ctx, query, id).Scan(
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

func (u *UserRepository) CreateUser(ctx context.Context, req dto.AuthRegisterRequest) (model.UserRegister, error) {
	query := `
INSERT INTO users (fullname, email, password)
VALUES ($1,$2,$3)
RETURNING id, fullname, email, password;
`

	rows, err := u.db.Query(ctx, query,
		req.FullName,
		req.Email,
		req.Password,
	)
	if err != nil {
		return model.UserRegister{}, err
	}

	newUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.UserRegister])
	if err != nil {
		return model.UserRegister{}, err
	}

	u.invalidateUserCache(ctx)
	return newUser, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user model.UserModel) error {
	query := `
UPDATE users SET
	fullname = $1,
	email = $2,
	password = $3,
	phone = $4,
	address = $5,
	picture = $6,
	role = $7,
	updated_at = now()
WHERE id = $8;
`

	_, err := u.db.Exec(ctx, query,
		user.FullName,
		user.Email,
		user.Password,
		user.Phone,
		user.Address,
		user.Picture,
		user.Role,
		user.Id,
	)
	if err != nil {
		return err
	}

	u.invalidateUserCache(ctx)
	return nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	u.invalidateUserCache(ctx)
	return nil
}
