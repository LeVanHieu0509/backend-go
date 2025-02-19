CREATE DATABASE IF NOT EXISTS shopdepgo
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- 1. ticket table
CREATE TABLE IF NOT EXISTS `shopdevgo`.`ticket` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
  `name` VARCHAR(50) NOT NULL COMMENT 'Ticket name',
  `desc` TEXT COMMENT 'Ticket description',
  `start_time` DATETIME NOT NULL COMMENT 'Ticket sale start time',
  `end_time` DATETIME NOT NULL COMMENT 'Ticket sale end time',
  `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'Ticket sale activity status',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last update time',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  PRIMARY KEY (`id`),
  KEY `idx_end_time` (`end_time`) COMMENT 'Very high query runtime',
  KEY `idx_start_time` (`start_time`) COMMENT 'Very high query runtime',
  KEY `idx_status` (`status`) COMMENT 'Very high query runtime'
) ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `shopdevgo`.`ticket_item` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
  `name` VARCHAR(50) NOT NULL COMMENT 'Ticket title',
  `description` TEXT COMMENT 'Ticket description',
  `stock_initial` INT(11) NOT NULL DEFAULT 0 COMMENT 'Initial stock quantity',
  `stock_available` INT(11) NOT NULL DEFAULT 0 COMMENT 'Current available stock',
  `is_stock_prepared` BOOLEAN NOT NULL DEFAULT 0 COMMENT 'Indicates if stock is prepared',
  `price_original` BIGINT(20) NOT NULL COMMENT 'Original ticket price',
  `price_flash` BIGINT(20) NOT NULL COMMENT 'Discounted price during flash sale',
  `sale_start_time` DATETIME NOT NULL COMMENT 'Flash sale start time',
  `sale_end_time` DATETIME NOT NULL COMMENT 'Flash sale end time',
  `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'Ticket status (e.g., activity id)',
  `activity_id` BIGINT(20) NOT NULL COMMENT 'ID of associated activity',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of the last update',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation timestamp',
  PRIMARY KEY (`id`),
  KEY `idx_end_time` (`sale_end_time`),
  KEY `idx_start_time` (`sale_start_time`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4
COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket details';

-- Insert into 'ticket' table
INSERT INTO `shopdevgo`.`ticket` (`name`, `desc`, `start_time`, `end_time`, `status`, `updated_at`, `created_at`)
VALUES
  ('Đặt Mở Bán Vé Ngày 12/12', 'Sự kiện mở bán vé đặc biệt cho ngày 12/12', '2024-12-12 00:00:00', '2024-12-12 23:59:59', 1, NOW(), NOW()),
  ('Đặt Mở Bán Vé Ngày 01/01', 'Sự kiện mở bán vé cho ngày đầu năm mới 01/01', '2025-01-01 00:00:00', '2025-01-01 23:59:59', 1, NOW(), NOW());

-- Insert data into 'ticket_item' table corresponding to each event in 'ticket' table
INSERT INTO `shopdevgo`.`ticket_item` (
  `name`, 
  `description`, 
  `stock_initial`, 
  `stock_available`, 
  `is_stock_prepared`, 
  `price_original`, 
  `price_flash`, 
  `sale_start_time`, 
  `sale_end_time`, 
  `status`, 
  `activity_id`, 
  `updated_at`, 
  `created_at`
)
VALUES
  ('Vé Sự Kiện 12/12', 'Vé phổ thông cho sự kiện ngày 12/12', 1000, 1000, 1, 100000, 80000, '2024-12-12 00:00:00', '2024-12-12 23:59:59', 1, 101, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Vé Sự Kiện 12/12', 'Vé VIP cho sự kiện ngày 12/12', 500, 500, 1, 200000, 160000, '2024-12-12 00:00:00', '2024-12-12 23:59:59', 1, 101, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Vé Sự Kiện 01/01', 'Vé phổ thông cho sự kiện ngày 01/01', 1000, 1000, 1, 300000, 240000, '2025-01-01 00:00:00', '2025-01-01 23:59:59', 1, 102, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('Vé Sự Kiện 01/01', 'Vé VIP cho sự kiện ngày 01/01', 500, 500, 1, 400000, 320000, '2025-01-01 00:00:00', '2025-01-01 23:59:59', 1, 102, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
