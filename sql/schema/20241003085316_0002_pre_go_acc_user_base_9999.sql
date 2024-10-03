-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_go_acc_user_base_9999` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `user_account` varchar(255) NOT NULL,
  `user_password` varchar(255) NOT NULL,
  `user_salt` varchar(55) NOT NULL, -- Sửa từ "varcha" thành "varchar(55)"
  `user_login_time` timestamp NULL DEFAULT NULL, -- Sửa từ "mestamp" thành "timestamp"
  `user_logout_time` timestamp NULL DEFAULT NULL,
  `user_login_ip` varchar(45) DEFAULT NULL,
  `user_created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `user_updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `unique_user_account` (`user_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4   COMMENT='pre_go_acc_user_base_9999';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`; 
-- +goose StatementEnd
