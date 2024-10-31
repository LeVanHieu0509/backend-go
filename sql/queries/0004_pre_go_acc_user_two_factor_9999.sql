-- Tạo một bản ghi mới cho phương thức xác thực hai yếu tố bằng email, đặt trạng thái là chưa kích hoạt (FALSE).
-- name: EnableTwoFactorTypeEmail :exec
INSERT INTO pre_go_acc_user_two_factor_9999 (user_id, two_factor_auth_type, two_factor_email, two_factor_auth_secret, two_factor_is_active, two_factor_created_at, two_factor_updated_at)
VALUES (?, ?, ?, 'OTP', FALSE, NOW(), NOW());

-- Vô hiệu hóa phương thức xác thực hai yếu tố cho người dùng bằng cách đặt trạng thái two_factor_is_active là FALSE.
-- name: DisableTwoFactor :exec
UPDATE pre_go_acc_user_two_factor_9999
SET two_factor_is_active = FALSE,
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- Cập nhật trạng thái của phương thức xác thực hai yếu tố thành TRUE nếu nó đang bị vô hiệu (FALSE)
-- name: UpdateTwoFactorStatus :exec
UPDATE pre_go_acc_user_two_factor_9999
SET two_factor_is_active = TRUE, two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = FALSE;

-- Kiểm tra xem phương thức xác thực hai yếu tố cụ thể có đang được kích hoạt cho người dùng hay không.
-- name: VerifyTwoFactor :one
SELECT COUNT(*)
FROM pre_go_acc_user_two_factor_9999
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = TRUE;

-- Lấy trạng thái hiện tại (active hoặc inactive) của phương thức xác thực hai yếu tố cho người dùng.
-- name: GetTwoFactorStatus :one
SELECT two_factor_is_active
FROM pre_go_acc_user_two_factor_9999
WHERE user_id = ? AND two_factor_auth_type = ?;

-- Kiểm tra xem có bất kỳ phương thức xác thực hai yếu tố nào đang được bật cho người dùng hay không.
-- name: IsTwoFactorEnabled :one
SELECT COUNT(*)
FROM pre_go_acc_user_two_factor_9999
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- Thêm hoặc cập nhật số điện thoại/email cho xác thực hai yếu tố, đặt trạng thái là active.
-- name: AddOrUpdatePhoneNumber :exec
INSERT INTO pre_go_acc_user_two_factor_9999 (user_id, two_factor_phone, two_factor_is_active)
VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE
    two_factor_phone = ?,
    two_factor_updated_at = NOW();

-- hêm hoặc cập nhật số điện thoại/email cho xác thực hai yếu tố, đặt trạng thái là active.
-- name: AddOrUpdateEmail :exec
INSERT INTO pre_go_acc_user_two_factor_9999 (user_id, two_factor_email, two_factor_is_active)
VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE
    two_factor_email = ?,
    two_factor_updated_at = NOW();

-- Lấy tất cả thông tin về các phương thức xác thực hai yếu tố của một người dùng cụ thể.
-- name: GetUserTwoFactorMethods :many
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret, two_factor_phone, two_factor_email,
       two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM pre_go_acc_user_two_factor_9999
WHERE user_id = ?;

-- Kích hoạt lại một phương thức xác thực hai yếu tố cho người dùng.
-- name: ReactivateTwoFactor :exec
UPDATE pre_go_acc_user_two_factor_9999
SET two_factor_is_active = TRUE,
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- Xóa phương thức xác thực hai yếu tố cho người dùng
-- name: RemoveTwoFactor :exec
DELETE FROM pre_go_acc_user_two_factor_9999
WHERE user_id = ? AND two_factor_auth_type = ?;

-- Đếm số lượng phương thức xác thực hai yếu tố đang được kích hoạt cho người dùng.
-- name: CountActiveTwoFactorMethods :one
SELECT COUNT(*)
FROM pre_go_acc_user_two_factor_9999
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- Lấy chi tiết của một phương thức xác thực hai yếu tố dựa trên two_factor_id.
-- name: GetTwoFactorMethodByID :one
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret, two_factor_phone, two_factor_email,
       two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM pre_go_acc_user_two_factor_9999
WHERE two_factor_id = ?;

-- Lấy chi tiết của một phương thức xác thực hai yếu tố dựa trên user_id và loại xác thực (two_factor_auth_type).
-- name: GetTwoFactorMethodByIDAndType :one
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret, two_factor_phone, two_factor_email,
       two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM pre_go_acc_user_two_factor_9999
WHERE user_id = ? AND two_factor_auth_type = ?;
