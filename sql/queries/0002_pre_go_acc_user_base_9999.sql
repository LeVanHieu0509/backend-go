-- name: GetOneUserInfo :one

SELECT user_id, user_account, user_password, user_salt
FROM `pre_go_acc_user_base_9999`
WHERE user_account = ?;

-- name: GetOneUserInfoAdmin :one
SELECT user_id, user_account, user_password, user_salt, user_login_time, user_logout_time,
user_created_at, user_updated_at
FROM `pre_go_acc_user_base_9999`
WHERE user_account = ?;

-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM pre_go_acc_user_base_9999
WHERE user_account = ?;

-- name: AddUserBase :execresult
INSERT INTO pre_go_acc_user_base_9999 (
    user_account, user_password, user_salt, user_created_at, user_updated_at
) VALUES (
    ?, ?, ?, NOW(), NOW()
);