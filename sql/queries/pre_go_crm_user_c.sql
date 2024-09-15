-- name: GetUserByEmailSQLC :one
SELECT usr_email, usr_id FROM `pre_go_crm_user_c` WHERE usr_email = ? LIMIT 1;

-- name: UpdateUserStatusByUserId :exec
UPDATE `pre_go_crm_user_c`
SET usr_status = ?,  -- Thay $2 bằng ?
    usr_updated_at = ?  -- Thay $3 bằng ?
WHERE usr_id = ?;  -- Thay $1 bằng ?