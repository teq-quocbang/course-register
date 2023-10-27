CREATE TABLE IF NOT EXISTS `register` (
  `account_id` BIGINT(20),
  `course_id` VARCHAR(50),
  `semester_id` VARCHAR(50),
  `class_id` VARCHAR(50),
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT,
  `is_canceled` BOOLEAN DEFAULT false,

  FOREIGN KEY (`account_id`) REFERENCES account(id),
  FOREIGN KEY (`course_id`) REFERENCES course(id),
  FOREIGN KEY (`semester_id`) REFERENCES semester(id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;