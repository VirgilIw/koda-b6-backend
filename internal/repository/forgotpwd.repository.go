package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/model"
)

type ForgotPwdRepository struct {
	db *pgxpool.Pool
}

func NewForgotPwdRepository(db *pgxpool.Pool) *ForgotPwdRepository {
	return &ForgotPwdRepository{
		db: db,
	}
}

func (f *ForgotPwdRepository) GetDataByEmailAndCode(ctx context.Context, req dto.ForgotPwdRequest) (model.ForgotPassword, error) {
	query := `SELECT id, email, code_otp, created_at
				FROM forgot_password
				WHERE email = $1 AND code_otp = $2`

	rows, err := f.db.Query(ctx, query, req.Email, req.CodeOtp)

	if err != nil {
		return model.ForgotPassword{}, err
	}

	forgotPwd, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ForgotPassword])

	if err != nil {
		return model.ForgotPassword{}, err
	}

	return forgotPwd, nil

}

func (f *ForgotPwdRepository) DeleteDataByCode(ctx context.Context, req dto.ForgotPwdRequest) error {
	query := `DELETE FROM forgot_password WHERE code_otp = $1; `

	_, err := f.db.Exec(ctx, query, req.CodeOtp)

	if err != nil {
		return err
	}

	return nil
}

func (f *ForgotPwdRepository) CreateForgotRequest(ctx context.Context, req dto.ForgotPwdRequest) (model.ForgotPassword, error) {
	query := `
		INSERT INTO forgot_password (email, code_otp)
		VALUES ($1, $2)
		RETURNING id, email, code_otp, created_at
	`

	rows, err := f.db.Query(ctx, query, req.Email, req.CodeOtp)

	if err != nil {
		return model.ForgotPassword{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ForgotPassword])

	if err != nil {
		return model.ForgotPassword{}, err
	}

	return data, nil
}
