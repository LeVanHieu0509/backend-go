-- name: GetUser :one
SELECT user_id, user_account, user_avatar, user_nickname, user_state, user_mobile, user_birthday, user_email, user_gender, user_is_authentication, created_at, updated_at
FROM `pre_go_acc_user_info_9999`
WHERE user_id = ? 
LIMIT 1;

-- name: GetUsers :many
SELECT user_id, user_account, user_avatar, user_nickname, user_state, user_mobile, user_birthday, user_email, user_gender, user_is_authentication, created_at, updated_at
FROM `pre_go_acc_user_info_9999`
WHERE user_id IN (?);

-- name: findUsers :many
SELECT *
FROM `pre_go_acc_user_info_9999`
WHERE user_account LIKE ? OR user_nickname LIKE ?;

-- name: listUsers :many
SELECT *
FROM `pre_go_acc_user_info_9999`
LIMIT ? OFFSET ?;

-- name: RemoveUser :exec
DELETE FROM `pre_go_acc_user_info_9999`
WHERE user_id = ?;

-- name: AddUserAutoUserId :execresult
INSERT INTO `pre_go_acc_user_info_9999` (
    user_account, user_avatar, user_nickname, user_state, user_mobile, user_birthday, user_email, user_gender, user_is_authentication
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: AddUserHaveUserId :execresult
INSERT INTO `pre_go_acc_user_info_9999` (
    user_id, user_account, user_avatar, user_nickname, user_state, user_mobile, user_birthday, user_email, user_gender, user_is_authentication
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: EditUserByUserId :execresult
UPDATE `pre_go_acc_user_info_9999`
SET user_nickname = ?, user_avatar = ?, user_mobile = ?, user_gender = ?, user_birthday = ?, user_email = ?, updated_at = NOW()
WHERE user_id = ? AND user_is_authentication = 1;