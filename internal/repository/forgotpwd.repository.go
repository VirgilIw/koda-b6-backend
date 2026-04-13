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

func (f *ForgotPwdRepository) GetDataByEmailAndCode(ctx context.Context, req dto.ForgotPwdRequest, codeOtp int) (model.ForgotPassword, error) {

	query := `
		SELECT id, email, code_otp, created_at
		FROM forgot_password
		WHERE email = $1 AND code_otp = $2
	`

	rows, err := f.db.Query(ctx, query, req.Email, codeOtp)

	if err != nil {
		return model.ForgotPassword{}, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ForgotPassword])

	if err != nil {
		return model.ForgotPassword{}, err
	}

	return data, nil
}

func (f *ForgotPwdRepository) DeleteDataByCode(ctx context.Context, req dto.ForgotPwdRequest, codeOtp int) error {

	query := `
		DELETE FROM forgot_password
		WHERE email = $1 AND code_otp = $2
	`

	_, err := f.db.Exec(ctx, query, req.Email, codeOtp)

	if err != nil {
		return err
	}

	return nil
}

func (f *ForgotPwdRepository) CreateForgotRequest(ctx context.Context, req dto.ForgotPwdRequest, codeOtp int) (model.ForgotPassword, error) {
	// otp pakai crypto generate random otp
	query := `
			INSERT INTO forgot_password (email, code_otp)
		VALUES ($1, $2)
		ON CONFLICT (email)
		DO UPDATE SET 
			code_otp = EXCLUDED.code_otp,
			created_at = NOW()
		RETURNING id, email, code_otp, created_at
	`

	rows, err := f.db.Query(ctx, query, req.Email, codeOtp)

	if err != nil {
		return model.ForgotPassword{}, err
	}
	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.ForgotPassword])

	if err != nil {
		return model.ForgotPassword{}, err
	}
	return data, nil
}
