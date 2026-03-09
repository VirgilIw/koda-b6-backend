package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]model.UserModel, error) {
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
FROM users;
	`

	rows, err := u.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[model.UserModel])
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

	err := u.conn.QueryRow(ctx, query, email).Scan(
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
