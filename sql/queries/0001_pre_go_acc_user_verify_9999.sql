-- name: GetValidOTP :one
SELECT verify_id_otp, verify_key_hash, verify_key, verify_id 
FROM `pre_go_acc_user_verify_9999`
WHERE verify_key_hash = ? AND is_verified = 0;

-- name: UpdateUserVerificationStatus :exec
UPDATE `pre_go_acc_user_verify_9999`
SET is_verified = 1,
    verify_updated_at = NOW() -- Đúng cú pháp cập nhật thời gian
WHERE verify_key_hash = ?; -- Bỏ dấu chấm thừa sau `verify_key_hash`

-- name: InsertOTPVerify :execresult
INSERT INTO `pre_go_acc_user_verify_9999` (
    verify_id_otp,
    verify_key,
    verify_key_hash,
    verify_type,
    is_verified,
    is_deleted,
    verify_created_at,
    verify_updated_at
) 
VALUES (?, ?, ?, ?, 0, 0, NOW(), NOW()); -- Đã đóng dấu ngoặc đúng cách

-- name: GetInfoOTP :one
SELECT verify_id, verify_id_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted, verify_created_at 
FROM `pre_go_acc_user_verify_9999`
WHERE verify_key_hash = ?;
