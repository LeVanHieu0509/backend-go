// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0001_pre_go_acc_user_verify_9999.sql

package database

import (
	"context"
	"database/sql"
)

const getInfoOTP = `-- name: GetInfoOTP :one

SELECT verify_id, verify_id_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted, verify_created_at 
FROM ` + "`" + `pre_go_acc_user_verify_9999` + "`" + `
WHERE verify_key_hash = ?
`

type GetInfoOTPRow struct {
	VerifyID        int32
	VerifyIDOtp     string
	VerifyKey       string
	VerifyKeyHash   string
	VerifyType      sql.NullInt32
	IsVerified      sql.NullInt32
	IsDeleted       sql.NullInt32
	VerifyCreatedAt sql.NullTime
}

// Đã đóng dấu ngoặc đúng cách
func (q *Queries) GetInfoOTP(ctx context.Context, verifyKeyHash string) (GetInfoOTPRow, error) {
	row := q.db.QueryRowContext(ctx, getInfoOTP, verifyKeyHash)
	var i GetInfoOTPRow
	err := row.Scan(
		&i.VerifyID,
		&i.VerifyIDOtp,
		&i.VerifyKey,
		&i.VerifyKeyHash,
		&i.VerifyType,
		&i.IsVerified,
		&i.IsDeleted,
		&i.VerifyCreatedAt,
	)
	return i, err
}

const getValidOTP = `-- name: GetValidOTP :one
SELECT verify_id_otp, verify_key_hash, verify_key, verify_id 
FROM ` + "`" + `pre_go_acc_user_verify_9999` + "`" + `
WHERE verify_key_hash = ? AND is_verified = 0
`

type GetValidOTPRow struct {
	VerifyIDOtp   string
	VerifyKeyHash string
	VerifyKey     string
	VerifyID      int32
}

func (q *Queries) GetValidOTP(ctx context.Context, verifyKeyHash string) (GetValidOTPRow, error) {
	row := q.db.QueryRowContext(ctx, getValidOTP, verifyKeyHash)
	var i GetValidOTPRow
	err := row.Scan(
		&i.VerifyIDOtp,
		&i.VerifyKeyHash,
		&i.VerifyKey,
		&i.VerifyID,
	)
	return i, err
}

const insertOTPVerify = `-- name: InsertOTPVerify :execresult

INSERT INTO ` + "`" + `pre_go_acc_user_verify_9999` + "`" + ` (
    verify_id_otp,
    verify_key,
    verify_key_hash,
    verify_type,
    is_verified,
    is_deleted,
    verify_created_at,
    verify_updated_at
) 
VALUES (?, ?, ?, ?, 0, 0, NOW(), NOW())
`

type InsertOTPVerifyParams struct {
	VerifyIDOtp   string
	VerifyKey     string
	VerifyKeyHash string
	VerifyType    sql.NullInt32
}

// Bỏ dấu chấm thừa sau `verify_key_hash`
func (q *Queries) InsertOTPVerify(ctx context.Context, arg InsertOTPVerifyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertOTPVerify,
		arg.VerifyIDOtp,
		arg.VerifyKey,
		arg.VerifyKeyHash,
		arg.VerifyType,
	)
}

const updateUserVerificationStatus = `-- name: UpdateUserVerificationStatus :exec
UPDATE ` + "`" + `pre_go_acc_user_verify_9999` + "`" + `
SET is_verified = 1,
    verify_updated_at = NOW() -- Đúng cú pháp cập nhật thời gian
WHERE verify_key_hash = ?
`

func (q *Queries) UpdateUserVerificationStatus(ctx context.Context, verifyKeyHash string) error {
	_, err := q.db.ExecContext(ctx, updateUserVerificationStatus, verifyKeyHash)
	return err
}