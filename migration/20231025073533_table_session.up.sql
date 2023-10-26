CREATE TABLE IF NOT EXISTS `session` (
  `id` BINARY(16) PRIMARY KEY,
  `account_id` BIGINT,
  `refresh_token` VARCHAR(500),
  `user_agent` VARCHAR(200),
  `client_ip`  VARCHAR(200),
  `expires_at` TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;