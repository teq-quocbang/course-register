CREATE TABLE IF NOT EXISTS `account_verify` (
  `account_id` BIGINT,
  `secret_code` VARCHAR(500),
  `is_used` BOOLEAN DEFAULT false,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;