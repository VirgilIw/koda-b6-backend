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
		if err == pgx.ErrNoRows {
			return model.UserModel{}, err
		}
		return model.UserModel{}, err
	}

	return user, nil
}

func (u *UserRepository) GetById(ctx context.Context, id int) (model.UserModel, error) {
	query := `
SELECT id, fullname, email, password, phone, address, picture, role, created_at, updated_at, deleted_at, lastlogin_at
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
	updated_at = $8
WHERE id = $9;
`

	_, err := u.db.Exec(ctx, query,
		user.FullName,
		user.Email,
		user.Password,
		user.Phone,
		user.Address,
		user.Picture,
		user.Role,
		time.Now(),
		user.Id,
	)

	if err != nil {
		return err
	}

	// Optional: invalidate cache
	u.rdb.Del(ctx, "users:all")

	return nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) CreateUser(ctx context.Context, req dto.AuthRegisterRequest) (model.UserRegister, error) {
	query := `
INSERT INTO users (fullname, email, password)
VALUES ($1,$2,$3)
Returning id, fullname, email, password
`
	// Simpan hasil QueryRow di variable row
	data, err := u.db.Query(ctx, query,
		req.FullName,
		req.Email,
		req.Password,
	)

	if err != nil {
		return model.UserRegister{}, err
	}

	var newUser model.UserRegister

	newUser, err = pgx.CollectOneRow(data, pgx.RowToStructByName[model.UserRegister])

	if err != nil {
		return model.UserRegister{}, err
	}
	// Hapus cache supaya GetUsers tidak menampilkan data lama
	u.rdb.Del(ctx, "users:all")

	return newUser, nil
}
