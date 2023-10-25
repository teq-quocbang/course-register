CREATE TABLE IF NOT EXISTS `account` (
    `id` BIGINT AUTO_INCREMENT,
    `username` VARCHAR(30) NOT NULL,
    `email` VARCHAR(200) NOT NULL,
    `hash_password` longblob NOT NULL,
    `is_verified` BOOLEAN DEFAULT false,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;