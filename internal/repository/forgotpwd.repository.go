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
