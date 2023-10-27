CREATE TABLE IF NOT EXISTS `course` (
  `id` VARCHAR(50),
  `semester_id` VARCHAR(50),
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` INT,

  PRIMARY KEY (`id`),
  FOREIGN KEY (`semester_id`) REFERENCES semester(id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;