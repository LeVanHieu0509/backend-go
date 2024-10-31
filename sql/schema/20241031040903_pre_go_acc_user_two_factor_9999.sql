-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS `pre_go_acc_user_two_factor_9999` (
    `two_factor_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, -- Khoá chính, tự động tăng
    `user_id` INT UNSIGNED NOT NULL, -- Mã người dùng liên quan đến tính năng xác thực hai yếu tố.
    `two_factor_auth_type` ENUM('SMS', 'EMAIL', 'APP') NOT NULL, -- Loại xác thực (SMS, EMAIL, APP), lưu trữ dưới dạng ENUM.
    `two_factor_auth_secret` VARCHAR(255) NOT NULL, -- Bí mật dùng để xác thực hai yếu tố (ví dụ như mã bảo mật hoặc key cho ứng dụng xác thực).
    `two_factor_phone` VARCHAR(20) NULL, -- Số điện thoại và email có thể sử dụng để gửi mã xác thực hai yếu tố.
    `two_factor_email` VARCHAR(255) NULL,
    `two_factor_is_active` BOOLEAN NOT NULL DEFAULT TRUE, -- Trạng thái hoạt động của xác thực hai yếu tố (BOOLEAN), mặc định là TRUE.
    `two_factor_created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Dấu thời gian lưu thời điểm tạo và thời điểm cập nhật bản ghi.
    `two_factor_updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX `idx_user_id` (`user_id`), -- Tạo chỉ mục trên cột user_id để tối ưu hoá các truy vấn liên quan đến người dùng.
    INDEX `idx_auth_type` (`two_factor_auth_type`) -- Tạo chỉ mục trên cột two_factor_auth_type để tối ưu hoá các truy vấn liên quan đến loại xác thực.
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='pre_go_acc_user_two_factor_9999'; -- Định nghĩa bảng sử dụng cơ chế lưu trữ InnoDB, hỗ trợ giao dịch và khóa (locking).


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_two_factor_9999`;
-- +goose StatementEnd
